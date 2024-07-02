import React, { useRef, useState } from "react";
import KeyboardReact from "react-simple-keyboard";
import "react-simple-keyboard/build/css/index.css";
import "./KeyBoard.css";
import Swal from "sweetalert2";
import { useTranslation } from "react-i18next";

const Keyboard = ({ onSubmit, onTouch }) => {
    const [layoutName, setLayoutName] = useState("default");
    const [userId, setUserId] = useState('')
    const { t } = useTranslation();
    const [error, setError] = useState(false)

    const onChange = (newInput) => {
        setError(false)
        setUserId(newInput)
    };

    const onKeyPress = (button) => {
        if (button === "{shift}" || button === "{lock}") {
            handleShift()
        }
    };
    const handleShift = () => {
        const newLayoutName = layoutName === "default" ? "shift" : "default";
        setLayoutName(newLayoutName);
    }

    const keyboardRef = useRef(null);

    const handleCancel = () => {
        Swal.close()
        setUserId('')
    }
    return (
        <div style={{ padding: '0 10px' }} onClick={onTouch}>
            <div className="message-box">
                <div>
                    <img src='/images/lock.png' alt="img33" width={40} />
                </div>
                <p>{t('Scan QR or enter user ID to login')}</p>
            </div>
            <div className="input-container">
                <input
                    className="custom-input"
                    placeholder={t("User ID")}
                    value={userId}
                    onChange={() => { }}
                />
                {error && <p className="please-enter">{t('Please enter user Id')}</p>}
                <button className="btn btn-primary swal-btn two" id="btn_vkeyboard_login"
                    onClick={() => {
                        userId ? onSubmit(userId) : setError(true)
                    }}
                >
                    {t('CONTINUE')}
                </button>
                <button className="btn btn-primary swal-btn two format-two" id="btn_vkeyboard_Cancel"
                    onClick={() => handleCancel()}
                >
                    {t('CANCEL')}
                </button>
            </div>

            <div cl='keyboard-container'>
                <KeyboardReact
                    // ref={keyboardRef}
                    layoutName={layoutName}
                    onChange={onChange}
                    onKeyPress={onKeyPress}
                />
            </div>
        </div>
    );
};

export default Keyboard;
