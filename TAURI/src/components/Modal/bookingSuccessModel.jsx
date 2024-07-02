import React, { useEffect } from "react";
import Swal from "sweetalert2";
import withReactContent from "sweetalert2-react-content";
import { useSelector } from "react-redux";
import { toast } from "react-toastify";
import i18n from "../../components/Layout/Language/i18n_react";
import { logoutUser } from "../../redux/actions/auth";
import store from "../../redux";
import { printTickit } from "../../redux/actions/bookings";

export const bookingSuccessModel = (bookingData, t) => {
    const selectedLanguage = i18n.language;
    const handlePrint = async (exSwal) => {
        const formatedData = {
            USER_NAME: bookingData.userData?.NAME,
            USER_ID: bookingData.userData?.SCHOOL_NO,
            BOOKING_NO: bookingData.bseqno,
            BOOKING_DATE: bookingData.date,
            ROOM_NO: bookingData.sector.ROOM_NAME_EN,
            FLOOR: bookingData.FLOOR_NAME_EN,
            SEAT_NO: bookingData.seatno,
            SEAT_TYPE: bookingData.sector.SECTOR_TYPE_NAME,
            EXTENSIONS: `0`,
            CHECKIN_TIME: bookingData.startTime,
            CHECKOUT_TIME: bookingData.endTime,
            QRCODE: bookingData.userData?.SCHOOL_NO,
            LANGUAGE: selectedLanguage
        }
        // dispatch(printTickit(formatedData))
        handleClose()
    }

    const handleClose = () => {
        store.dispatch(logoutUser());
        Swal?.close();
    }

    const timers = store.getState().setupReducer.timers;

    const MySwal = withReactContent(Swal);
    MySwal.fire({
        showConfirmButton: false,
        showCancelButton: false,
        width: 800,
        timer: 1000 * timers.bookingsDetails,
        html: (
            <div>
                <div className="seat-confirm">
                    <div className="thank-icon">
                        <img src="/images/check2.png" alt="confirmation-icon"></img>
                    </div>
                    <div className="thank-comment" style={{ paddingLeft: '80px' }}>
                        <h4>{t("SUCCERSS")}</h4>
                        <p>{t("Check your seat details in your reciept")}.</p>
                    </div>
                </div>
                <div className="reciept-detail">
                    <div className="reciept-half" style={{ borderBottom: '1px solid #b6b6b6' }}>
                        <h6>{t("Seat No.")}</h6>
                        <p>
                            {bookingData.seatno}
                        </p>
                        <h6>{t("Student ID")}</h6>
                        <p>{bookingData.userData?.SCHOOL_NO}</p>
                        <h6>{t("Start Time")}</h6>
                        <p>{bookingData.startTime}</p>
                        <h6>{t("Booking Id")}</h6>
                        <p>{bookingData.bseqno}</p>

                    </div>
                    <div className="reciept-half" style={{ borderBottom: '1px solid #b6b6b6' }}>
                        <h6>{t("Room")}</h6>
                        <p> {bookingData.sector.ROOM_NAME} </p>
                        <h6>{t("User Name")}</h6>
                        <p> {bookingData.userData?.NAME}</p>
                        <h6>{t("End Time")}</h6>
                        <p> {bookingData.endTime} </p>

                        <h6>{t("Booking Date")}</h6>
                        <p> {bookingData.date} </p>
                    </div>
                </div>
                <div className="button-container">
                    <button className="btn btn-primary swal-btn" id="btn1" onClick={() => handlePrint(Swal)}>
                        {t("PRINT")}
                    </button>
                    <button className="btn btn-secondary swal-btn" id="btn2" onClick={() => handleClose(Swal)}>
                        {t("OK")}
                    </button>
                </div>
            </div>
        ),
    }).then(() => {
        handleClose()
    })
}