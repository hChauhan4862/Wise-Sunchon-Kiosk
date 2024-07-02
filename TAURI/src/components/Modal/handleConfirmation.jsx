import React, { useEffect } from "react";
import Swal from "sweetalert2";
import withReactContent from "sweetalert2-react-content";
import moment from 'moment';
import { logoutUser } from "../../redux/actions/auth";
import { renderToString } from "react-dom/server";
import store from "../../redux";
import * as api from '../../apis';
import i18n from "../../components/Layout/Language/i18n_react";
import { hcAlert } from "./hcAlert";

export const handleConfirmation = ({
    navigate,
    uiSelectedSector,
    toast,
    t,
}) => {
    const userData = store.getState().authReducer.userData
    const timers = store.getState().setupReducer.timers;
    const hcLiveSync = store.getState().setupReducer.hcLiveSync;
    if (!userData || !userData?.user) {
        return
    }
    console.log(timers, 'timers')

    const booking = userData?.booking || [];
    const user = userData.user;
    // Get the First Booking of each type
    let carelBooking = booking.filter(function (el) { return el.Typeno == 22 });
    carelBooking = (carelBooking.length > 0) ? carelBooking[0] : null;

    let groupBooking = booking.filter(function (el) { return el.Typeno == 10 });
    groupBooking = (groupBooking.length > 0) ? groupBooking[0] : null;

    let normalBooking = booking.filter(function (el) { return el.Typeno != 22 && el.Typeno != 10 });
    normalBooking = (normalBooking.length > 0) ? normalBooking[0] : null;



    const data = {
        navigate,
        userID: user.SCHOOL_NO,
        userName: user.NAME,
        bookings: [groupBooking, carelBooking, normalBooking],
        allActiveBookings: booking,
        uiSelectedSector,
        toast,
        timers,
        t, // Translation function
        kioskConfig: hcLiveSync.KioskConfig,
    };

    bookingOptions(data, 0);
}

const bookingOptions = (data, seq) => {
    console.log(data, seq, "bookingOptions")
    if (seq >= data.bookings.length) {
        console.log("No Booking Found")
        data.navigate(`/floor/${data.uiSelectedSector}`);
        Swal?.close();
        return;
    };
    if (!data.bookings[seq]) { bookingOptions(data, seq + 1); return; }

    const booking = data.bookings[seq];
    const typeName = booking.Typeno == 22 ? "Carel" : booking.Typeno == 10 ? "Group" : "";

    const MySwal = withReactContent(Swal);
    Swal?.close();
    MySwal.fire({
        imageUrl: '/images/info.png',
        imageAlt: 'Info',
        title: `Choose options for your ${typeName} booking in ${i18n.language === "ko" ? booking.SectorName : booking.SectorNameEn} room`,
        showConfirmButton: false,
        width: 1000,
        imageWidth: 120,
        timer: data.timers.popup * 1000,
        allowOutsideClick: false,
        html: (
            <div className="button-container">
                {
                    booking.ALLOW_ReturnYN == "Y" && data.kioskConfig.Returnable && (
                        <button id="BTN_TO_DIABLE_1" className="btn btn-primary swal-btn" onClick={(e) => bookingCancel(data, seq, e)}>
                            {data.t("RETURN")}
                        </button>
                    )
                }
                {
                    booking.ALLOW_CancelYN == "Y" && data.kioskConfig.Returnable && (
                        <button id="BTN_TO_DIABLE_2" className="btn btn-primary swal-btn" onClick={(e) => bookingCancel(data, seq, e)}>
                            CANCEL
                        </button>
                    )
                }
                {
                    booking.ALLOW_MoveYn == "Y" && data.kioskConfig.Movable && (
                        <button id="BTN_TO_DIABLE_3" className="btn btn-primary swal-btn" onClick={(e) => seatChange(data, seq, e)}>
                            {data.t("CHANGE")}
                        </button>
                    )
                }
                {
                    booking.ALLOW_ExtendYn == "Y" && data.kioskConfig.Extendable && (
                        <button id="BTN_TO_DIABLE_4" className="btn  btn-primary swal-btn" onClick={(e) => bookingExtend(data, seq, e)}>
                            {data.t("EXTEND")}
                        </button>
                    )
                }
                {
                    true && (
                        <button id="BTN_TO_DIABLE_5" className="btn btn-primary swal-btn" onClick={(e) => myBookingModal(data, seq, e)}>
                            {data.t("BOOKINGS")}
                        </button>
                    )
                }
                {
                    booking.ALLOW_ConfirmYN == "Y" && (
                        <button id="BTN_TO_DIABLE_6" className="btn btn-primary swal-btn" onClick={(e) => bookingConfirm(data, seq, e)} >
                            {/* {data.t('YES')} */}
                            ASSIGN
                        </button>
                    )
                }
                {
                    (seq == data.bookings.length - 1 || !data.kioskConfig.Assignable) && ( // This is the last booking Means user has normal booking
                        <button id="BTN_TO_DIABLE_7" className="btn btn-secondary swal-btn format-two" onClick={() => logout(data)}>
                            {data.t("LOGOUT")}
                        </button>
                    )
                }
                {
                    (seq !== data.bookings.length - 1) && data.kioskConfig.Assignable && ( // This is not the last booking Means user has more booking
                        <button id="BTN_TO_DIABLE_8" className="btn btn-primary swal-btn format-two" onClick={() => bookingOptions(data, seq + 1)}>
                            {data.t('CONTINUE')}
                        </button>
                    )
                }
            </div>
        ),
    }).then((result) => {
        if (result.dismiss === Swal.DismissReason.timer) {
            logout(data);
        }
    });
}




/********************************************* BUTTON ACTIONS *********************************************/

const seatChange = (data, seq) => {
    var sector = data.bookings[seq].Sectorno;
    data.navigate(`/floor/${sector}?type=change&bseqno=${data.bookings[seq].Bseqno}`);
    Swal?.close()
}

const logout = (data) => {
    store.dispatch(logoutUser())
    data?.navigate("/")
    Swal?.close()
}

/**********************************************   CANCEL BOOKING  **************************************/
const bookingCancel = (data, seq, e) => {
    for (let i = 1; i <= 10; i++) {
        if (document.getElementById(`BTN_TO_DIABLE_${i}`)) {
            document.getElementById(`BTN_TO_DIABLE_${i}`).disabled = true;
        }
    }
    (async () => {
        try {
            const json = {
                bseqno: parseInt(data.bookings[seq].Bseqno),
            }
            const res = await api.returnSeat(json)
            if (res.data.error) {
                hcAlert("Oops!", res.data.error, true)
            } else {
                hcAlert("Success!", null, true)
            }
        } catch (error) {
            hcAlert("Oops!", "Something went wrong!", true)
        }
    })();
}
/**********************************************   CANCEL BOOKING  **************************************/








/**********************************************   EXTEND BOOKING  **************************************/
const bookingExtend = (data, seq, e) => {
    for (let i = 1; i <= 10; i++) {
        if (document.getElementById(`BTN_TO_DIABLE_${i}`)) {
            document.getElementById(`BTN_TO_DIABLE_${i}`).disabled = true;
        }
    }
    (async () => {
        try {
            let ExtendTill = data.bookings[seq].ExtendTill;
            if (ExtendTill == null) {
                hcAlert("Oops!", "Booking cannot be extended", true)
                return;
            }
            const now = new Date();
            ExtendTill = new Date(ExtendTill);
            ExtendTill = new Date(ExtendTill.getTime() + (now.getTimezoneOffset()) * 60000); // Make it Equel to local time Zone
            const newTime = `${String(ExtendTill.getHours()).padStart(2, '0')}${String(ExtendTill.getMinutes()).padStart(2, '0')}`;
            const minutesDiff = Math.floor(((ExtendTill - now) / 1000) / 60);

            const json = {
                bseqno: parseInt(data.bookings[seq].Bseqno),
                minute: minutesDiff,
                timeEnd: newTime
            }
            const res = await api.extendSeat(json)
            console.log(res.data)
            if (res.data.error) {
                hcAlert("Oops!", res.data.error, true)
            } else {
                hcAlert("Success!", null, true)
            }
        } catch (error) {
            console.log(error)
            hcAlert("Oops!", "Something went wrong!", true)
        }
    })();
}
/**********************************************   EXTEND BOOKING  **************************************/









/**********************************************   CONFIRM BOOKING  **************************************/
const bookingConfirm = (data, seq, e) => {
    for (let i = 1; i <= 10; i++) {
        if (document.getElementById(`BTN_TO_DIABLE_${i}`)) {
            document.getElementById(`BTN_TO_DIABLE_${i}`).disabled = true;
        }
    }
    (async () => {
        try {
            const json = {
                bseqno: parseInt(data.bookings[seq].Bseqno)
            }
            const res = await api.confirmSeat(json)
            if (res.data.error) {
                hcAlert("Oops!", res.data.error, true)
            } else {
                hcAlert("Success!", null, true)
            }
        } catch (error) {
            hcAlert("Oops!", "Something went wrong!", true)
        }
    })();
}
/**********************************************   CONFIRM BOOKING  **************************************/





/**********************************************      MY BOOKING    **************************************/
const myBookingModal = (data, seq) => {
    const booking = data.allActiveBookings || [];
    const carelBooking = booking.filter(function (el) { return el.Typeno == 22 });
    const groupBooking = booking.filter(function (el) { return el.Typeno == 10 });
    const normalBooking = booking.filter(function (el) { return el.Typeno != 22 && el.Typeno != 10 });

    const MySwal = withReactContent(Swal);
    Swal?.close();
    MySwal.fire({
        showConfirmButton: false,
        width: 900,
        imageHeight: 900,
        imageWidth: 1000,
        timer: 1000 * (data?.timers?.popup || 10),
        allowOutsideClick: false,
        html: (
            <>
                <div className="seat-confirm two">
                    <div className="thank-comment">
                        <h4>{data.t("YOUR BOOKINGS")}</h4>
                    </div>
                </div>
                <div className="reciept-detail two">
                    <div className="cards-container">
                        {groupBooking.length && (
                            <>
                                <h4>{data.t("Group Study Booking")}</h4>
                                {
                                    groupBooking.map((d, idx) => {
                                        return (
                                            <div>
                                                <div className="custom-card">
                                                    <div className="nested-card-2">
                                                        <div className="row-container">
                                                            <p>{data.t("Booking Id")}</p>
                                                            <p>{d.Bseqno}</p>
                                                        </div>
                                                        <div className="row-container">
                                                            <p>{data.t("User ID")}</p>
                                                            <p style={{ 'blue': '#acacac' }} >{data.userID}</p>
                                                            {
                                                                d.MEMBERS?.split(",").map((item, idx) => (
                                                                    (idx % 2 == 0) && (<p key={idx}>{item}</p>)
                                                                ))
                                                            }
                                                        </div>

                                                        <div className="row-container">
                                                            <p>{data.t("User Name")}</p>
                                                            <p style={{ 'blue': '#acacac' }} >{data.userName}</p>
                                                            {
                                                                d.MEMBERS.split(",").map((item, idx) => (
                                                                    (idx % 2 == 1) && (<p key={idx}>{item}</p>)
                                                                ))
                                                            }
                                                        </div >
                                                    </div>
                                                    {/* <div className="row-container">
                                                        <p>
                                                            <span style={{ width: 500, display: "inline-block" }}>{data.t("Start")}</span>
                                                            <span style={{ width: 500, display: "inline-block" }}>{data.t("End")}</span>
                                                        </p>
                                                        <p>
                                                            <span style={{ width: 500, display: "inline-block" }}>{d.Usestart}</span>
                                                            <span style={{ width: 500, display: "inline-block" }}>{d.Useexpire}</span>
                                                        </p>
                                                    </div> */}
                                                    <div className="row-container">
                                                        <p>{data.t("Start")}</p>
                                                        <p>{d.Usestart}</p>
                                                    </div>
                                                    <div className="row-container">
                                                        <p>{data.t("End")}</p>
                                                        <p>{d.Useexpire}</p>
                                                    </div>
                                                </div >
                                            </div>
                                        )
                                    })
                                }
                            </>
                        ) || <div className="no-booking"></div>}
                        {
                            normalBooking.length && (
                                <div>
                                    <h4>{data.t("Reading Room Booking")}</h4>
                                    {
                                        normalBooking.map((d, idx) => {
                                            return (
                                                <>
                                                    <div className="custom-card">
                                                        <div className="nested-card-2">
                                                            <div className="row-container">
                                                                <p>{data.t("Booking Id")}</p>
                                                                <p>{d.Bseqno}</p>
                                                            </div>
                                                            <div className="row-container">
                                                                <p>{data.t("Seat Details")}</p>
                                                                {
                                                                    i18n.language === "ko" && (
                                                                        <p>{d.SeatName} {d.SeatVname}</p>
                                                                    )
                                                                }
                                                                {
                                                                    i18n.language !== "ko" && (
                                                                        <p>{d.SeatNameEn} {d.SeatVnameEn}</p>
                                                                    )
                                                                }
                                                            </div>
                                                            <div className="row-container">
                                                                <p>{data.t("Start")}</p>
                                                                <p>{d.Usestart}</p>
                                                            </div>
                                                            <div className="row-container">
                                                                <p>{data.t("End")}</p>
                                                                <p>{d.Useexpire}</p>
                                                            </div>
                                                        </div>
                                                    </div>
                                                </>
                                            )
                                        }
                                        )}
                                </div>
                            ) || <div className="no-booking"></div>
                        }
                        {
                            carelBooking.length && (
                                <div>
                                    <h4>{data.t("Carrel Room Booking")}</h4>
                                    {
                                        carelBooking.map((d, idx) => {
                                            return (
                                                <>
                                                    <div className="custom-card">
                                                        <div className="nested-card-2">
                                                            <div className="row-container">
                                                                <p>{data.t("Booking Id")}</p>
                                                                <p>{d.Bseqno}</p>
                                                            </div>
                                                            <div className="row-container">
                                                                <p>{data.t("Room Details")}</p>
                                                                <p>{d?.SectorName} {d?.FloorName}</p>
                                                            </div>
                                                            <div className="row-container">
                                                                <p>{data.t("Start Date")}</p>
                                                                <p>{moment(d.Usestart).format("D MMM")}</p>
                                                            </div>
                                                            <div className="row-container">
                                                                <p>{data.t("End Date")}</p>
                                                                <p>{moment(d.Useexpire).format("D MMM")}</p>
                                                            </div>
                                                        </div>
                                                    </div>
                                                </>
                                            )
                                        }
                                        )}
                                </div>
                            ) || <div className="no-booking"></div>
                        }
                    </div >
                </div >
                <div className="button-container">
                    <button className="btn btn-secondary swal-btn" onClick={() => bookingOptions(data, seq)}>
                        {data.t("Back")}
                    </button>
                    <button className="btn btn-primary swal-btn format-two" onClick={() => logout(data,)}>
                        {data.t("LOGOUT")}
                    </button>
                </div>
            </>
        ),
    }).then((result) => {
        if (result.dismiss === Swal.DismissReason.timer) {
            logout(data);
        }
    });
};
/**********************************************     MY BOOKING     **************************************/