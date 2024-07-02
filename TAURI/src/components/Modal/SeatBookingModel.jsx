import React, { useEffect } from "react";
import Swal from "sweetalert2";
import withReactContent from "sweetalert2-react-content";
import { logoutUser } from "../../redux/actions/auth";
import store from "../../redux";
// import { useTranslation as t } from "react-i18next";
import { printTickit } from "../../redux/actions/bookings";
import { useSelector } from "react-redux";
import { toast } from "react-toastify";

export const SeatBookingModel = (navigate, bookingData, userData, t, timers, dispatch, selectedLanguage) => {
  const handlePrint = async (exSwal) => {
    const formatedData = {
      USER_NAME: userData.userName,
      USER_ID: userData.userID,
      BOOKING_NO: bookingData.BookingId,
      BOOKING_DATE: bookingData.BookingDate,
      ROOM_NO: bookingData.roomNumber,
      FLOOR: bookingData.floorName,
      SEAT_NO: bookingData.seatNo,
      SEAT_TYPE: bookingData.seatCategoryName,
      EXTENSIONS: `${bookingData.extension_count}/${bookingData.totalExtention}`,
      CHECKIN_TIME: bookingData.use_start_time,
      CHECKOUT_TIME: bookingData.use_end_time,
      QRCODE: userData.barcode,
      LANGUAGE: selectedLanguage
    }
    // dispatch(printTickit(formatedData))
    handleClose()
  }

  const handleClose = () => {
    store.dispatch(logoutUser());
    Swal?.close();
  }

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
              {bookingData.seatNo}
            </p>
            <h6>{t("Student ID")}</h6>
            <p>{bookingData.user_id}</p>
            <h6>{t("Start Time")}</h6>
            <p>{bookingData.use_start_time.slice(0, 5)}</p>
            <h6>{t("Booking Id")}</h6>
            <p>{bookingData.BookingId}</p>

          </div>
          <div className="reciept-half" style={{ borderBottom: '1px solid #b6b6b6' }}>
            <h6>{t("Room")}</h6>
            <p> {bookingData.roomName} </p>
            <h6>{t("User Name")}</h6>
            <p> {bookingData.user_name}</p>
            <h6>{t("End Time")}</h6>
            <p> {bookingData.use_end_time.slice(0, 5)} </p>

            <h6>{t("Booking Date")}</h6>
            <p> {bookingData.BookingDate} </p>
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