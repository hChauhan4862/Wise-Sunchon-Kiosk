import React, { useState, useEffect, useMemo } from "react";
import { useTranslation } from "react-i18next";
import { Link } from "react-router-dom";
import { useLocation } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";
import { logoutUser } from "../../../redux/actions/auth";
import { useNavigate } from 'react-router-dom'
import { askLogin } from "../../Modal/askLogin";
import { getUserByQr, normalLogin } from "../../../redux/actions/auth";
import store from '../../../redux'
import Swal from "sweetalert2";
import withReactContent from "sweetalert2-react-content";
import Keyboard from "../../KeyBoard/KeyBoard";
import { emit, listen } from '@tauri-apps/api/event'


const Header = () => {
  const path = useLocation().pathname;
  const hcLiveSync = useSelector((state) => state.setupReducer.hcLiveSync);
  const [pathName, setPathName] = useState(path);
  const { t } = useTranslation();
  const userData = useSelector(state => state.authReducer.userData)

  const [eventUnsub, setEventUnsub] = useState(() => async () => { });

  const dispatch = useDispatch()
  const timers = useSelector(state => state.setupReducer.timers)
  const socket = listen

  const handleSocketResponse = async (data) => {
    console.log("Event Emmited :: ", data)
    const userData = store.getState().authReducer.userData
    if (userData === null || userData === undefined) {
      dispatch(normalLogin(data))
    } else {
      dispatch(getUserByQr(data))
    }
  }


  useEffect(() => {
    if (hcLiveSync?.KioskConfig?.allowUseRfidDevice) {
      const unlisten = listen("SERIAL_EVENT", (e) => {
        handleSocketResponse(e.payload.message)
      });
      return () => {
        unlisten.then(f => f());
      };
    }

  }, [socket, hcLiveSync])

  const handleLogout = () => {
    dispatch(logoutUser())
  }

  const handleLoginSubmit = (userId) => {
    dispatch(normalLogin(userId))
  }

  const handleKeyboardTouch = () => {
    Swal.increaseTimer(timers?.keyboard * 1000 - Swal.getTimerLeft())
  }

  const handleLogin = () => {
    if (hcLiveSync?.KioskConfig?.allowUseVirtualKeyboard) {
      const MySwal = withReactContent(Swal);
      MySwal.fire({
        showConfirmButton: false,
        width: 1300,
        timer: timers?.keyboard * 1000,
        imageWidth: 110,
        html: <Keyboard
          onSubmit={(userId) => {
            document.getElementById('btn_vkeyboard_login').disabled = true
            document.getElementById('btn_vkeyboard_Cancel').disabled = true
            handleLoginSubmit(userId)
          } }
          onTouch={handleKeyboardTouch}
        />,
        allowOutsideClick: false,
      });
    } else {
      askLogin(t)
    }
  }

  const randerIcon = useMemo(() => {
    if (path === '/') {
      return <img src="/images/logo-main.png" alt="Header Logo"></img>
    } else {
      return <img src="/images/logo-white.png" alt="Header Logo"></img>
    }
  }, [userData, path])


  const getTimer = useMemo(() => {
    if (path !== '/') {
      return <div className="time-stamp-container" id="time" ></div>
    } else {
      <></>
    }

  }, [path])


  return (
    <div>
      <div id="top-header">
        <div className="header-logo">
          {randerIcon}
        </div>

        <div className="header-profile">
          {getTimer}

          {userData ? (
            <div className="profile-stat">
              <img src="/images/profile.png" alt="Flag Turkey"></img>{" "}
              {userData.user.NAME}
            </div>
          ) : (
            <div className="profile-stat" style={{ display: "none" }}></div>
          )}
          <div
            className={pathName === "/" ? "logout-button" : "logout-button2"}
          >
            {userData ? (
              <Link onClick={handleLogout} to='/'>
                {t("Logout")}{" "}
                <img src="/images/logout.png" alt="Login Button"></img>
              </Link>
            ) : (
              <Link
                // to="/"
                onClick={handleLogin}
                className="login-btn"
              >
                {t("Login")}{" "}
                <img src="/images/logout.png" alt="Logout Button"></img>
              </Link>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Header;
