import React, { useEffect } from "react";
import Swal from "sweetalert2";
import withReactContent from "sweetalert2-react-content";
import { logoutUser } from "../../redux/actions/auth";
import store from "../../redux";

export const askLogin = (t) => {
  const MySwal = withReactContent(Swal);
  return MySwal.fire({
    title: <p>{t("Please Login to continue")}</p>,
    width: 600,
    timer: 1000,
    html: <div>{t("Scan your Qr or barcode for login")}</div>,
    showConfirmButton: false,
    imageUrl: '/images/lock.png',
    allowOutsideClick: false
  });
};
