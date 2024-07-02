import React, { useEffect, useState, useTransition } from "react";
import { Link } from "react-router-dom";
import NoticeBoard from "../../components/NoticeBoard/NoticeBoard";
import { useDispatch, useSelector } from "react-redux";
import { askLogin } from "../../components/Modal/askLogin";
import { useNavigate } from "react-router-dom";
import { handleConfirmation } from "../../components/Modal/handleConfirmation";
import store from "../../redux";
import { useTranslation } from "react-i18next";
import i18n from "../../components/Layout/Language/i18n_react";
import Swal from "sweetalert2";
import withReactContent from "sweetalert2-react-content";
import Keyboard from "../../components/KeyBoard/KeyBoard";
import { normalLogin } from "../../redux/actions/auth";
import { toast } from "react-toastify";

const Home = () => {
  const { t } = useTranslation();
  const navigate = useNavigate();
  const dispatch = useDispatch();

  const hcLiveSync = useSelector((state) => state.setupReducer.hcLiveSync);

  const userData = useSelector((state) => state.authReducer.userData);
  const timers = useSelector((state) => state.setupReducer.timers);

  const [navigatesectorID, setnavigatesectorID] = useState(null);

  useEffect(() => {
    if (hcLiveSync?.softAvailability?.length && !navigatesectorID) {
      const id = hcLiveSync.softAvailability[0]?.SECTORS[0]?.SECTORNO;
      setnavigatesectorID(id);
    }
  }, [hcLiveSync]);

  useEffect(() => {
    if (userData) {
      let data = {
        navigate,
        uiSelectedSector: navigatesectorID,
        toast,
        t,
      };
      handleConfirmation(data);
    }
  }, [userData]);

  const handleLoginSubmit = (userId) => {
    const MySwal = withReactContent(Swal);
    const errorAlert = (message) => {
      MySwal.fire({
        title: t("Opps!"),
        showConfirmButton: false,
        width: 600,
        imageUrl: "/images/cancel.png",
        timer: timers.error * 1000,
        imageWidth: 110,
        html: <div>{t(message)}</div>,
      });
    };
    const closeModal = () => {
      if (Swal) {
        Swal.close();
      }
    };
    dispatch(normalLogin(userId, closeModal, errorAlert));
  };

  const handleKeyboardTouch = () => {
    Swal.increaseTimer(timers?.keyboard * 1000 - Swal.getTimerLeft());
  };

  const handleLogin = (FIRST_SECTOR_NO) => {
    setnavigatesectorID(FIRST_SECTOR_NO);

    if (true) {
      const MySwal = withReactContent(Swal);
      MySwal.fire({
        showConfirmButton: false,
        width: 1300,
        timer: timers?.keyboard * 1000,
        imageWidth: 110,
        html: (
          <Keyboard
            onTouch={handleKeyboardTouch}
            onSubmit={handleLoginSubmit}
          />
        ),
        allowOutsideClick: false,
      });
    } else {
      askLogin(t);
    }
  };

  return (
    <div>
      <div className="main-container">
        <img
          className="main-logo"
          src="images/logo-main.png"
          alt="sunchon logo main"
        ></img>
        <h1>{t("Welcome to Sunchon National University")}</h1>
        <div className="front-selc-btn">
          {hcLiveSync?.softAvailability?.map((sector, idx) => {
            return (
              <Link
                key={sector.SECTOR_TYPE_NO}
                className={"option-btn btn-color-" + (idx + 1)}
                onClick={() => handleLogin(sector.SECTORS[0]?.SECTORNO)}
              >
                <div style={{height: '180px'}}>
                  <h5> {i18n.language === "ko" ? sector.SECTOR_TYPE_NAME : sector.SECTOR_TYPE_NAME_EN} </h5>
                  <h6>{" "} {sector.SECTORS.length} Area</h6>
                </div>
                <div className="bt-dtl">
                  {t("Available")} - {sector?.TOTAL_CNT - (sector?.USE_CNT + sector?.FIX_CNT)} {/* / {" "} {sector?.TOTAL_CNT} */}
                </div>
              </Link>
            );
          })}
        </div>
        <h4>{t("Please select the desired floor and obtain a seat")}</h4>
        {true && <NoticeBoard />}
      </div>
    </div>
  );
};

export default Home;
