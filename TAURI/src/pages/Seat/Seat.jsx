import React, { useEffect, useState } from "react";
import { Link, useLocation, useNavigate, useParams } from "react-router-dom";
import BookingModal from "../../Component/Modals/BookingModal";
import { useTranslation } from "react-i18next";
import { data } from "../../dummyData";
import { useDispatch, useSelector } from "react-redux";
import { fetchFloors } from "../../Redux/actions";
import ROOT from "../../utils/root.json";
import axios from "axios";
import jwtDecode from "jwt-decode";
import swal from "sweetalert";
import Loader from "../../Component/subcomponent/Loader";
import { printHandler } from "../../utils/print";


const Seat = () => {
  const { t } = useTranslation();
  const navigate = useNavigate();

  const param = useParams();
  const zoneId = param.id;
  const location = useLocation();
  const queryParams = new URLSearchParams(location.search);
  const [openFloorId, setOpenFloorId] = useState(floorId);
  



  return (
    <div id="main-floors">
      {/* <BookingModal
        bookingData={bookingData}
        show={modalShow}
        onHide={closeModal}
        handleRelocate={handleRelocate}
      /> */}
      <div className="floor-map">
        <div className="breadcrumb">
          <div className="layers-nav">{t("Floor")}{floorInfo?.filter((floor) => floor?.floorId == openFloorId)[0].floorName}{" "}»{" "}
            {
              floorInfo
                ?.filter((floor) => floor.floorId == openFloorId)[0]
                .zones.filter((zone) => zone.zoneId == zoneId)[0].roomName
            }
          </div>
          <div className="important-tag">
            <ul>
              <li>
                <div className="seat-available"></div>
                {t("Available")}
              </li>
              <li>
                <div className="seat-unavialable"></div>
                {t("Unavailable")}
              </li>
            </ul>
          </div>
        </div>

        {isloading ? (
          <Loader />
        ) : (
          <div
            className="floor-image"
            style={{
              position: "relative",
            }}
          >
            <a>
              <img
                src={ROOT.HOST + "kiosk/displayzonemap/" + zoneId}
                useMap="#image_map"
              />

              {!closed ? (
                <div className="closed-message">
                  <h2>{t("library closed")}</h2>
                </div>
              ) : (
                zoneWiseSeats &&
                zoneWiseSeats?.map((item, index) => {
                  const seatBooked = Loading
                    ? seatStatus
                        .filter((seat) => seat.seatId == item.seatId)
                        .map((it) => it.status)
                    : "";

                  return (
                    <div
                      className={
                        item.isFixed
                          ? "box seat seatFixed"
                          : `box seat ${seatBooked[0]} `
                      }
                      key={item?.seatId}
                      style={{
                        top: `${item.coords?.split(",")[1]}px`,
                        left: `${item.coords?.split(",")[0]}px`,
                        width: `${
                          item.coords?.split(",")[2] -
                          item.coords?.split(",")[0]
                        }px`,
                        height: `${
                          item.coords?.split(",")[3] -
                          item.coords?.split(",")[1]
                        }px`,
                      }}
                      onClick={() => {
                        !userDecode?.LastBooking
                          ? createBooking(item.seatId, item.isFixed)
                          : seatChangeHandler(item.seatId, item.isFixed);
                      }}
                    >
                      {item.seatNumber}
                    </div>
                  );
                })
              )}
            </a>
          </div>
        )}
      </div>
      <div className="floor-list">
        {floorInfo &&
          floorInfo.map((floor) => (
            <div key={floor.floorId}>
              <h1>
                <Link
                  className={openFloorId == floor.floorId ? "active" : ""}
                  to={`/floor/${floor.floorId}`}
                  onClick={() => handleClick(floor.floorId)}
                >
                  {t(floor.floorName)}
                </Link>
              </h1>
              <div
                className={`floor-zones ${
                  openFloorId == floor.floorId ? "open" : "closed"
                }`}
              >
                <ul>
                  {floor.zones.map((zone) => {
                    const zonePercentage = Loading
                      ? seatPercentage &&
                        seatPercentage.zonePercentage &&
                        seatPercentage.zonePercentage.filter(
                          (percentageItem) =>
                            percentageItem.zoneId === zone.zoneId
                        )[0]
                      : "";
                    return (
                      <li key={zone.zoneId}>
                        <Link
                          to={`/seatMap/${zone.zoneId}?floorId=${openFloorId}`}
                        >
                          {t(zone.roomName)}
                          <span>
                            {t("Available")}{" "}
                            {zonePercentage ? zonePercentage.percentage : ""}
                          </span>
                        </Link>
                      </li>
                    );
                  })}
                </ul>
              </div>
            </div>
          ))}
      </div>
    </div>
  );
};

export default Seat;
