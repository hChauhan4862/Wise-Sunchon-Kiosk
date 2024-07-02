import React, { useState, useEffect, useRef } from "react";
import { Link, useParams, useNavigate } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";
import { useTranslation } from "react-i18next";
import { TransformWrapper, TransformComponent, MiniMap } from "react-zoom-pan-pinch";
import moment from "moment";
import Swal from "sweetalert2";

import { bookingSuccessModel } from "../../components/Modal/bookingSuccessModel";
import { logoutUser } from '../../redux/actions/auth'
import { SessionInterval, SessionTimer } from '../../components/SessionTimer/sessionLogoutTimer';

import { actionSeatsList, actionSeatsMapLoad } from "../../redux/actions/seats";
import { hcAlert } from "../../components/Modal/hcAlert";
import * as api from '../../apis'
import i18n from "../../components/Layout/Language/i18n_react";

const Floor = () => {
    const timers = useSelector(state => state.setupReducer.timers)
    const hcLiveSync = useSelector((state) => state.setupReducer.hcLiveSync);
    const seatsList = useSelector((state) => state.seatsReducer.seatsListBySector);
    const userData = useSelector((state) => state.authReducer.userData);
    const mapUrls = useSelector((state) => state.seatsReducer.seatsMapUrls);


    const [mapOriginalHeight, setMapOriginalHeight] = useState(null);
    const [mapOriginalWidth, setMapOriginalWidth] = useState(null);

    const [isBooking, setisBooking] = useState(false);

    const dispatch = useDispatch()
    const navigate = useNavigate();
    const { t } = useTranslation()

    const param = useParams();
    const sectorID = param.id


    function getQueryParams(url) {
        const params = {};
        const paramArray = url.slice(url.indexOf('?') + 1).split('&')
        paramArray.forEach(param => {
            const [key, value] = param.split('=')
            params[key] = value;
        });
        return params
    }
    const queryParams = getQueryParams(window.location.href)

    const [selectedSectorId, setSelectedSectorId] = useState(sectorID)
    const [selectedSector, setSelectedSector] = useState(null)
    const [selectedSectorType, setSelectedSectorType] = useState(null)
    const [isRoomClosed, setIsRoomClosed] = useState(false)

    const [bookingType, setbookingType] = useState(-1)
    const [bseqno, setBseqno] = useState(-1)

    const handleLogout = () => {
        dispatch(logoutUser())
        Swal?.close()
    }

    useEffect(() => {
        setbookingType(queryParams?.type)
        setBseqno(queryParams?.bseqno)
    }, [])

    useEffect(() => {
        console.log('seatsList', seatsList)
    }, [seatsList])

    useEffect(() => {
        if (hcLiveSync?.softAvailability?.length && !selectedSectorId) {
            const id = hcLiveSync.softAvailability[0]?.SECTORS[0]?.SECTORNO;
            setSelectedSectorId(id);
        }

        // Uncomment this if you want to check for seats list update Only if the seats are changed detectable
        // if (seatsList) {
        //     hcLiveSync?.softAvailability?.forEach(item => {
        //         item?.SECTORS?.forEach(SECTOR => {
        //             if (SECTOR.SECTORNO == selectedSectorId) {
        //                 let free = SECTOR?.TOTAL_CNT - SECTOR?.USE_CNT - SECTOR?.FIX_CNT
        //                 let free2 = seatsList?.Seats.filter(item => (item?.status == 1 || item?.status == 2) && item?.USECNT == 0).length
        //                 if (free != free2) {
        //                     console.log('dispatching', free, free2)
        //                     dispatch(actionSeatsList(selectedSectorId))
        //                 }
        //             }
        //         })
        //     })
        // }
        dispatch(actionSeatsList(selectedSectorId))
    }, [hcLiveSync]);


    useEffect(() => {
        dispatch(actionSeatsList(selectedSectorId))
        if (!selectedSectorId) {
            setSelectedSector(null)
            setIsRoomClosed(false)
            setSelectedSectorType(null)
            return
        }
        let SECTOR = null
        let isClosed = false
        hcLiveSync?.softAvailability?.forEach(item => {
            item?.SECTORS?.forEach(sector => {
                if (sector.SECTORNO == selectedSectorId) {
                    SECTOR = sector
                    if (selectedSector?.SECTOR_IMAGE != sector.SECTOR_IMAGE) {
                        setMapOriginalWidth(null);
                        setMapOriginalHeight(null);
                    }
                    setSelectedSectorType(item.SECTOR_TYPE_NO)
                }
            })
        })
        if (SECTOR) {
            const currentTime = moment((moment().format('YYYY-MM-DD HH:mm:ss')), 'YYYY-MM-DD HH:mm:ss')
            const startTimeMoment = moment(SECTOR.ROOM_OPENTIME, 'YYYY-MM-DD HH:mm:ss')
            const closeTimeMoment = moment(SECTOR.ROOM_CLOSETIME, 'YYYY-MM-DD HH:mm:ss')
            isClosed = !currentTime.isBetween(startTimeMoment, closeTimeMoment)
        }
        setSelectedSector(SECTOR)
        setIsRoomClosed(isClosed)
        if (!mapUrls || !mapUrls[SECTOR?.SECTOR_IMAGE]) {
            dispatch(actionSeatsMapLoad(SECTOR?.SECTOR_IMAGE));
        }
        // setIsRoomClosed(false)

    }, [selectedSectorId])

    const handleSeatSelect = async (seat) => {
        if (!(seat.USECNT == '0' && (seat.status == '1' || seat.status == '2'))) { return }

        if (isBooking) { return; }
        setisBooking(true);

        const seatId = seat?.seatno
        try {
            if (bookingType === 'change') {
                const data = {
                    bseqno: parseInt(bseqno),
                    new_seat_no: parseInt(seatId)
                }
                try {
                    const res = await api.changeSeat(data)
                    console.log(res.data);
                    if (res.data.error == false) {
                        hcAlert("Success", res.data.result, true)
                    }
                    else {
                        hcAlert("Opps!", res.data.result, true)
                    }
                } catch (error) {
                    console.log(error)
                    hcAlert("Opps!", "Something went wrong", true)
                }
            } else {
                const close = moment(selectedSector.ROOM_CLOSETIME, 'YYYY-MM-DD HH:mm:ss')
                const now = moment(moment().format('YYYY-MM-DD HH:mm:ss'), 'YYYY-MM-DD HH:mm:ss')
                const maxAllowed = now.add(seatsList?.ROOM?.USE_MIN, 'minutes')

                const endTime = (close.isBefore(maxAllowed) ? close : maxAllowed).format('HHmm')

                const startTime = moment().format('HHmm')
                const data = {
                    "userID": userData.user.SCHOOL_NO,
                    "seatNo": parseInt(seatId),
                    "timeStart": startTime,
                    "timeEnd": endTime,
                }

                try {
                    const res = await api.seatBooking(data)
                    console.log(res.data);
                    if (res.data.error == false) {
                        let newBSEQNO = res.data?.result?.bseqno
                        let bookingData = {
                            bseqno: newBSEQNO,
                            seatno: seatId,
                            sector: selectedSector,
                            startTime,
                            endTime,
                            date: moment().format('YYYY-MM-DD'),
                            userData: userData.user
                        }
                        bookingSuccessModel(bookingData, t)
                        // hcAlert("Success", res.data.result, true)
                    }
                    else {
                        hcAlert("Opps!", res.data.result, true)
                    }
                } catch (error) {
                    console.log(error)
                    hcAlert("Opps!", "Something went wrong", true)
                }
            }
        } catch (error) {
            if (error.response) {
                errorAlert(error.response.data.message, true)
            } else {
                errorAlert(error.message, true)
            }
        }
    }

    useEffect(() => {
        if (!userData) {
            navigate('/')
        }
        if (userData) {
            SessionInterval(handleLogout)
        }
    }, [userData])

    const mapRef = useRef(null);

    useEffect(() => {
        if (mapRef.current) {
            mapRef.current.resetTransform()
        }
    }, [selectedSectorId])

    const imageMeasurements = (image) => {
        if (!image) return;
        setMapOriginalWidth(image.naturalWidth);
        setMapOriginalHeight(image.naturalHeight);
    };

    const renderSeats = () => {
        return <>
            {
                seatsList?.Seats?.map((seat, index) => {
                    let use = (seat.USECNT == '0' && (seat.status == '1' || seat.status == '2')) ? '' : 'use';
                    let shape = (seat.ICONTYPE == '2' || seat.ICONTYPE == '3' || seat.ICONTYPE == '4' || seat.ICONTYPE == '5' || seat.ICONTYPE == '6' || seat.ICONTYPE == '7') ? 'hexa' : 'square';
                    let square_thin = (seat.ICONTYPE == '8') ? 'square_thin' : '';
                    let handicap = (seat.status == '9') ? 'handicap' : '';
                    return <>
                        <div key={seat.seatno}
                            className={`area_seat_btn ${shape} ${square_thin} ${handicap} ${use}`}
                            style={{ position: 'absolute', top: `${seat.POSY}px`, left: `${seat.POSX}px` }}>
                            <a onClick={() => { handleSeatSelect(seat) }} style={{ width: `${seat.POSW}px`, height: `${seat.POSH}px`, lineHeight: `${seat.POSH}px` }} className={`type${seat.ICONTYPE} ${use}`}>
                                <span className="txt_seat">
                                    {seat.vname}
                                </span>
                                {
                                    seat.XFULL != null && <span className="graph_sm" style={{ width: `${seat.XFULL}%` }}></span>
                                }
                            </a>
                        </div>
                        {
                            (use == 'use') && (
                                <img src="/images/seats/SeatIcons.png" style={{ position: 'absolute', top: `${seat.MANY}px`, left: `${seat.MANX}px` }} />
                            )
                        }
                    </>
                })
            }
        </>
    }

    return (
        <div id="main-floors">
            <div className="floor-map">
                <div className="breadcrumb">
                    <div className="layers-nav">
                        {i18n.language === "ko" ? selectedSector?.FLOOR_NAME : selectedSector?.FLOOR_NAME_EN}
                        <div className="zone-name-container">
                            <img src="/images/arrow-right.png" alt="arrow-right" />
                            {selectedSector?.MAPLABEL}
                            <img src="/images/arrow-right.png" alt="arrow-right" />
                            {i18n.language === "ko" ? selectedSector?.SECTOR_NAME : selectedSector?.SECTOR_NAME_EN}

                        </div>
                    </div>
                    <div className="important-tag">
                        {
                            true && (
                                <ul>
                                    <li>
                                        <div className="hcSeat-Laptop"></div> Laptop
                                    </li>
                                    <li>
                                        <div className="hcSeat-available"></div> {t("Available")}
                                    </li>
                                    <li>
                                        <div className="hcSeat-using"></div> In Use
                                    </li>
                                    <li>
                                        <div className="hcSeat-handicap"></div> Disabled
                                    </li>
                                </ul>
                            )
                        }
                    </div>
                </div>

                <TransformWrapper disablePadding={true} maxScale={2} ref={mapRef} disabled={isRoomClosed} style={{ position: "relative", }}>
                    <div style={{ position: "absolute", top: 198, right: 699, zIndex: 1000, }} >
                        <MiniMap><img src={(mapUrls != null) ? mapUrls[selectedSector?.SECTOR_IMAGE] : ""} /></MiniMap>
                    </div>
                    <TransformComponent >
                        {
                            (
                                <div className={`floor-image`} style={{ position: "relative", }} >
                                    {
                                        !moment((moment().format('YYYY-MM-DD HH:mm:ss')), 'YYYY-MM-DD HH:mm:ss').isBetween(moment(selectedSector?.ROOM_OPENTIME, 'YYYY-MM-DD HH:mm:ss'), moment(selectedSector?.ROOM_CLOSETIME, 'YYYY-MM-DD HH:mm:ss')) && (
                                            <div className="closed-message-container">
                                                <div>
                                                    <div>
                                                        <img src="/images/info.png"
                                                            alt="opps"
                                                            className="error-img"
                                                        />
                                                    </div>
                                                    <h5>{t('Ohh!')}</h5>
                                                    <p>{t('This room is closed')}</p>
                                                </div>
                                            </div>
                                        )
                                    }
                                    <img
                                        id="floor-map-image"
                                        onLoad={(e) => imageMeasurements(e.target)}
                                        // src = {api.getAPIBaseUrl() + "/uploads"+ selectedSector?.SECTOR_IMAGE}
                                        src={(mapUrls != null) ? mapUrls[selectedSector?.SECTOR_IMAGE] : ""}
                                    />

                                    {
                                        (mapOriginalHeight && mapOriginalWidth && (seatsList?.Seats?.length == 0 || seatsList?.Seats[0].sectorno == selectedSectorId))
                                        && (
                                            <div
                                                className="scaled-seats"
                                                style={{
                                                    transform: `scale(${1200 / mapOriginalWidth}, ${780 / mapOriginalHeight})`,
                                                    position: "fixed",
                                                    top: 0,
                                                    left: 0,
                                                }}
                                            >
                                                {renderSeats()}
                                            </div>
                                        )
                                    }
                                </div>
                            )
                        }
                    </TransformComponent>
                </TransformWrapper>

            </div>
            <div className="floor-list">
                {
                    hcLiveSync?.softAvailability?.map(sectorTypes => {
                        return (
                            <>
                                <div key={sectorTypes.SECTOR_TYPE_NO} style={{ overflow: "hidden" }}>
                                    <h1>
                                        <Link className={selectedSectorType == sectorTypes.SECTOR_TYPE_NO ? "active" : ""} onClick={() => setSelectedSectorType(sectorTypes.SECTOR_TYPE_NO)}>
                                            {i18n.language === "ko" ? sectorTypes.SECTOR_TYPE_NAME : sectorTypes.SECTOR_TYPE_NAME_EN}
                                        </Link>
                                    </h1>
                                </div>
                                <div className={`floor-zones ${selectedSectorType == sectorTypes.SECTOR_TYPE_NO ? "open" : "closed"}`}>
                                    <ul>
                                        {
                                            sectorTypes?.SECTORS?.map(sector => {
                                                return (
                                                    <li key={sector.SECTORNO} className={selectedSectorId == sector.SECTORNO ? 'selected' : ''}>
                                                        <Link onClick={() => setSelectedSectorId(sector.SECTORNO)} >
                                                            {sector.MAPLABEL}
                                                            {
                                                                // Condition for the room is closed
                                                                !moment((moment().format('YYYY-MM-DD HH:mm:ss')), 'YYYY-MM-DD HH:mm:ss').isBetween(moment(sector.ROOM_OPENTIME, 'YYYY-MM-DD HH:mm:ss'), moment(sector.ROOM_CLOSETIME, 'YYYY-MM-DD HH:mm:ss')) && (
                                                                    <span style={{backgroundColor:"#bf2e2e"}}>Closed</span>
                                                                ) || <span>{t("Available")} {sector.TOTAL_CNT - sector.USE_CNT - sector.FIX_CNT}</span>
                                                            }
                                                        </Link>
                                                    </li>
                                                )
                                            })
                                        }
                                    </ul>
                                </div>
                            </>
                        )

                    })
                }
            </div>
        </div>
    );
};

export default Floor
