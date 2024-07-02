import React, { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import Carousel from 'react-bootstrap/Carousel';
import { useSelector } from "react-redux";
import  i18n from "../../components/Layout/Language/i18n_react";

const NoticeBoard = () => {
    const { t } = useTranslation();
    const [notices, setNotices] = useState([]);
    const hcLiveSync = useSelector((state) => state.setupReducer.hcLiveSync);
    // console.log(notices);

    const [currentNoticeIndex, setCurrentNoticeIndex] = useState(0);
    const [isAnimating, setIsAnimating] = useState(false);

    useEffect(() => {
        if (hcLiveSync?.notice) {
            setNotices(hcLiveSync.notice);
        }
    }, [hcLiveSync]);

    // Function to move to the next notice
    const moveToNextNotice = () => {
        setIsAnimating(true);
        setTimeout(() => {
            setCurrentNoticeIndex((prevIndex) => (prevIndex + 1) % notices.length);
            setIsAnimating(false);
        }, 500);
    }

    useEffect(() => {
        if (notices && notices.length > 0) {
            const carouselTimeout = setTimeout(moveToNextNotice, 10000);
            return () => clearTimeout(carouselTimeout);
        }
    }, [currentNoticeIndex]);

    return (
        <div className="notice-board">
            <div className="notice-title">{t("Notice")}</div>

            <div className={`notice-main ${isAnimating ? 'animating' : ''}`}>
                <b>{t('Publish Date')}: {notices && notices[currentNoticeIndex]?.InsertTime}</b>
                <br />
                <p>{notices && i18n.language === "ko" && notices[currentNoticeIndex]?.Contents}</p>
                <p>{notices && i18n.language !== "ko" && notices[currentNoticeIndex]?.EnContents}</p>
            </div>
        </div>
    );
};

export default NoticeBoard;