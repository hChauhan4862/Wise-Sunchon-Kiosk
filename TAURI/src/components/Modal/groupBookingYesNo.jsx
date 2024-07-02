import React, { useEffect } from "react";
import Swal from "sweetalert2";
import withReactContent from "sweetalert2-react-content";
import {
  logoutUser,
  setUserForQR,
  groupStudyMembers,
} from "../../redux/actions/auth";
import {
  cunfirmGroupStudyBooking,
  returnGroupStudyBooking,
  returnSeatBooking,
  extendSeatBooking,
  cunfirmCarrelBooking,
  returnCarrelBooking
} from "../../redux/actions/bookings";
import { SeatBookingModel } from "./bookingSuccessModel";
import store from "../../redux";
import { renderToString } from "react-dom/server";
import moment from 'moment';
export const handleConfirmation = ({
  navigate,
  userId,
  groupBookingId,
  carrolBookingId,
  ReadingRoomBookingId,
  allowExtendReadingRoom,
  allowChangeReadingRoomSeat,
  floorID,
  roomId,
  zoneId,
  selectedSectorId,
  t,
  toast,
  carrelConfirmation
}) => {
  const kioskConfig = store.getState().setupReducer.kioskConfig
  const timers = store.getState().setupReducer.timers;
  const data = {
    navigate,
    userId,
    groupBookingId,
    carrolBookingId,
    ReadingRoomBookingId,
    allowExtendReadingRoom,
    allowChangeReadingRoomSeat,
    floorID,
    roomId,
    zoneId,
    selectedSectorId,
    t,
    kioskConfig,
    toast,
    carrelConfirmation
  }

  if (kioskConfig.allowTicketing) {
    askGroupBookingConfirm(data)
  } else {
    const MySwal = withReactContent(Swal);
    MySwal.fire({
      title: data.t('Opps!'),
      showConfirmButton: false,
      width: 600,
      imageUrl: "/images/cancel.png",
      timer: timers.error * 1000,
      imageWidth: 110,
      html: <div>{data.t('Tickting is not allowed!')}</div>,
      allowOutsideClick: false,
    }).then((result) => {
      if (result.dismiss === Swal.DismissReason.timer) {
        logout(data);
      }
    })
  }
}

const myBookingModal = (data, exSwal, prevMdl) => {
  const userData = store.getState().authReducer.userData;
  const timers = store.getState().setupReducer.timers;

  if (exSwal) {
    exSwal.close()
  }

  const handleBack = (data, xSwal) => {
    if (prevMdl === "GROUP") {
      askGroupBookingConfirm(data);
    }
    else if (prevMdl === 'CARREL') {
      askCarrolBookingConfirm(data);
    }
    else if (prevMdl === "SEAT") {
      askSeatBookingConfirm(data);
    }
  };
  const MySwal = withReactContent(Swal);
  MySwal.fire({
    showConfirmButton: false,
    width: 900,
    imageHeight: 900,
    imageWidth: 1000,
    timer: 1000 * timers.popup,
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
            {data.groupBookingId && (
              <>
                <h4>{data.t("Group Study Booking")}</h4>
                <div className="custom-card">

                  <div className="nested-card-2">
                    <div className="row-container">
                      <p>{data.t("Booking Id")}</p>
                      <p>{data.groupBookingId}</p>
                    </div>
                    <div className="row-container">
                      <p>{data.t("User Name")}</p>
                      {userData.LastGroupStudyBooking.member_user_name
                        .split("|")
                        .map((item, idx) => (
                          <p
                            style={{
                              color: item === userData.LastGroupStudyBooking.user_name ?
                                'blue' : '#acacac'
                            }}
                            key={idx}>{item}</p>
                        ))}
                    </div>
                    <div className="row-container">
                      <p>{data.t("User Id")}</p>
                      {userData.LastGroupStudyBooking.member_user_id
                        .split("|")
                        .map((item, idx) => (
                          <p style={{
                            color: item === userData.LastGroupStudyBooking.user_id ? 'blue' : '#acacac'
                          }}
                            key={idx}>{item}</p>
                        ))}
                    </div>

                    <div className="row-container">
                      <p>{data.t("Start")}</p>
                      <p>{userData.LastGroupStudyBooking.bookingStartTime.slice(0, 5)}</p>
                    </div>
                    <div className="row-container">
                      <p>{data.t("End")}</p>
                      <p>{userData.LastGroupStudyBooking.bookingEndTime.slice(0, 5)}</p>
                    </div>
                  </div>
                </div>
              </>
            )}
            {data.ReadingRoomBookingId && (
              <>
                <h4>{data.t("Reading Room Booking")}</h4>
                <div className="custom-card">
                  <div className="nested-card-2">
                    <div className="row-container">
                      <p>{data.t("Booking Id")}</p>
                      <p>{data.ReadingRoomBookingId}</p>
                    </div>
                    <div className="row-container">
                      <p>{data.t("Seat Details")}</p>
                      <p>{userData?.LastBooking.seat_category.name} {userData?.LastBooking.seatNo}</p>
                    </div>
                    <div className="row-container">
                      <p>{data.t("Start")}</p>
                      <p>{userData?.LastBooking.use_start_time.slice(0, 5)}</p>
                    </div>
                    <div className="row-container">
                      <p>{data.t("End")}</p>
                      <p>{userData?.LastBooking.use_end_time.slice(0, 5)}</p>
                    </div>
                  </div>
                </div>
              </>
            )}
            {data.carrolBookingId && (
              <>
                <h4>{data.t("Carrel Room Booking")}</h4>
                <div className="custom-card">

                  <div className="nested-card-2">
                    <div className="row-container">
                      <p>{data.t("Booking Id")}</p>
                      <p>{data.carrolBookingId}</p>
                    </div>
                    <div className="row-container">
                      <p>{data.t("Room Details")}</p>
                      <p>{userData?.lastCarrolBooking.carol_study_room.roomName} {userData?.lastCarrolBooking.floors.floorName}</p>
                    </div>
                    <div className="row-container">
                      <p>{data.t("Start Date")}</p>
                      <p>{moment(userData?.lastCarrolBooking.bookingDate).format("D MMM")}</p>
                    </div>
                    <div className="row-container">
                      <p>{data.t("End Date")}</p>
                      <p>{moment(userData?.lastCarrolBooking.bookingEndDate).format("D MMM")}</p>
                    </div>
                  </div>
                </div>
              </>
            )}
          </div>
        </div>
        <div className="button-container">
          <button className="btn btn-secondary swal-btn" id="btn2" onClick={() => handleBack(data, Swal)}>
            {data.t("Back")}
          </button>
          <button className="btn btn-primary swal-btn format-two" id="btn1" onClick={() => logout(data, Swal)}>
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

const logout = (data, exSwal) => {
  const timers = store.getState().setupReducer.timers;
  if (exSwal) {
    exSwal.close();
  }
  // const MySwal = withReactContent(Swal);
  // MySwal.fire({
  //   title: data.t("Success"),
  //   html: data.t("Logout successfully"),
  //   showConfirmButton: false,
  //   width: 600,
  //   imageUrl: "/images/check.png",
  //   timer: timers.success * 1000,
  //   imageWidth: 110,
  //   allowOutsideClick: false,
  // }).then((result) => {
  //   if (result.dismiss === Swal.DismissReason.timer) {
  store.dispatch(logoutUser())
  data?.navigate("/")
  exSwal?.close()
  // data.toast.success(data.t('Logout successful'))
  //   }
  // });
};

const groupBookingConfirmModal = (data, exSwal) => {
  const timers = store.getState().setupReducer.timers;
  if (exSwal) {
    exSwal.close();
  }
  const confirmBooking = (data, exSwal, bookingId) => {
    if (exSwal) {
      exSwal.close();
    }
    const MySwal = withReactContent(Swal);

    const successAlert = () => {
      MySwal.fire({
        title: data.t("Success"),
        html: data.t("Group study booking cunfirmed successfully"),
        showConfirmButton: false,
        width: 600,
        imageUrl: "/images/check.png",
        timer: timers.success * 1000,
        imageWidth: 110,
        allowOutsideClick: false,
      }).then((result) => {
        if (result.dismiss === Swal.DismissReason.timer) {
          logout(data);
        }
      });
    };

    const errorAlert = (message) => {
      MySwal.fire({
        title: data.t('Opps!'),
        showConfirmButton: false,
        width: 600,
        imageUrl: "/images/cancel.png",
        timer: timers.error * 1000,
        imageWidth: 110,
        html: <div>{data.t(message)}</div>,
        allowOutsideClick: false,
      }).then((result) => {
        if (result.dismiss === Swal.DismissReason.timer) {
          logout(data);
        }
      });
    };

    const finalData = store
      .getState()
      .authReducer.members.filter((item) => item.status === 1)
      .map((item) => item.id);

    const bookingdata = {
      bookingId: bookingId,
      userId: finalData,
    };
    store.dispatch(
      cunfirmGroupStudyBooking(bookingdata, successAlert, errorAlert)
    );
  };

  const MySwal = withReactContent(Swal);
  const userData = store.getState().authReducer.userData;
  let timerInterval = null;
  const memberIDs =
    userData?.LastGroupStudyBooking?.member_user_id?.split("|") || [];
  const memberNames =
    userData?.LastGroupStudyBooking?.member_user_name?.split("|") || [];
  let memberConfirm =
    userData?.LastGroupStudyBooking?.confirmed_user_ids?.split("|") || [];
  const memberReturn =
    userData?.LastGroupStudyBooking?.return_user_ids?.split("|") || [];
  let bookingId = null;
  if (userData.LastGroupStudyBooking) {
    bookingId = userData.LastGroupStudyBooking.bookingId;
  }

  let memberS = [];
  for (let i = 0; i < memberIDs.length; i++) {
    memberS.push({
      id: memberIDs[i],
      name: memberNames[i],
      status:
        memberConfirm.includes(memberIDs[i]) || memberIDs[i] == userData?.userID
          ? 1
          : memberReturn.includes(memberIDs[i])
            ? -1
            : 0,
    });
  }
  userData?.userID in memberConfirm || memberConfirm.push(userData?.userID);
  MySwal.fire({
    showConfirmButton: false,
    html: (
      <div id="confirmPopup">
        <h3>{data.t("Do you want to confirm your group booking?")}</h3>
      </div>
    ),
    timer: 1000 * timers.popup,
    timerProgressBar: true,
    width: 900,
    allowOutsideClick: false,
    didOpen: () => {
      const b = Swal.getHtmlContainer().querySelector("#confirmPopup");
      store.dispatch(setUserForQR(-1));
      store.dispatch(groupStudyMembers(memberS));
      timerInterval = setInterval(() => {
        const QRmember = store.getState().authReducer.userId || null;
        if (QRmember !== -1) {
          store.dispatch(setUserForQR(-1));
          if (QRmember === null) {
            data.toast.error(data.t("User Not Found"))
            console.log("User Not Found");
          } else {
            let member = null;
            for (let i = 0; i < memberS.length; i++) {
              if (memberS[i].id === QRmember) {
                console.log("Status", memberS[i].status);
                member = memberS[i];
                if (member.status === 0) {
                  memberS[i].status = 1;
                  memberConfirm.push(member.id);
                  store.dispatch(groupStudyMembers(memberS));
                  Swal.increaseTimer(10 * 1000 - Swal.getTimerLeft())
                } else if (member.status === 1) {
                  data.toast.error(data.t("You already confirmed"))
                  console.log("User Already Confirm");
                } else if (member.status === -1) {
                  data.toast.error(data.t("You returned this reservation"))
                  console.log("User Already Return this Reservation");
                }
                break;
              }
            }
            if (memberIDs.length === memberConfirm.length) {
              confirmBooking(data, Swal, bookingId);
            }
            if (member === null) {
              data.toast.error(data.t("You are not part of this Group Study"))
              console.log("You are not part of this Group Study");
            }
          }
        }
        const myHtml = (
          <div>
            <div className="message-box">
              <p style={{ fontWeight: 600, textAlign: 'left' }}>{data.t("Scan your QR / BARCODE to cunfirm booking")}</p>
              <div>
                <img src='/images/scanner.png' alt="img33" width={40} />
              </div>
            </div>

            <div className="custom-table">
              <div className="table-row">
                <p>{data.t("ID")}</p>
                <p>{data.t('Name')}</p>
                <p>{data.t('Status')}</p>
              </div>
              {memberS.map((member) => (

                <div className="table-row">
                  <p>{member.id}</p>
                  <p>{member.name}</p>
                  <p>
                    {member.status === 0
                      ? (<img src="/images/circle.png" alt="confirmation-icon"
                        width={30}
                      />)
                      : member.status === 1
                        ? <img src="/images/check2.png" alt="confirmation-icon5"
                          width={40}
                        />
                        : <img src="/images/cancel.png" alt="confirmation-icon2"
                          width={30}
                        />
                    }
                  </p>
                </div>
              ))}
            </div>
          </div>
        );
        b.innerHTML = renderToString(myHtml)
      }, 100);
    },
    willClose: () => {
      clearInterval(timerInterval);
    },
  }).then((result) => {
    if (result.dismiss === Swal.DismissReason.timer) {
      confirmBooking(data, Swal, bookingId);
    }
  })
}

const askGroupBookingConfirm = (data, exSwal) => {
  const timers = store.getState().setupReducer.timers;
  if (exSwal) {
    exSwal.close();
  }
  const userData = store.getState().authReducer.userData;

  const returnBooking = (data, exSwal) => {
    const MySwal = withReactContent(Swal);

    const successAlert = () => {
      MySwal.fire({
        title: data.t("Success"),
        html: data.t("Group study booking Returned Successfully"),
        showConfirmButton: false,
        width: 600,
        imageUrl: "/images/check.png",
        timer: timers.success * 1000,
        imageWidth: 110,
        allowOutsideClick: false,
      }).then((result) => {
        if (result.dismiss === Swal.DismissReason.timer) {
          logout(data);
        }
      });
    };
    const errorAlert = (message) => {
      MySwal.fire({
        title: data.t('Opps!'),
        showConfirmButton: false,
        width: 600,
        imageUrl: "/images/cancel.png",
        timer: timers.error * 1000,
        imageWidth: 110,
        html: <div>{data.t(message)}</div>,
        allowOutsideClick: false,
      }).then((result) => {
        if (result.dismiss === Swal.DismissReason.timer) {
          logout(data);
        }
      })
    };

    const returndata = {
      bookingId: userData.LastGroupStudyBooking.bookingId,
      userId: userData.userID,
    };
    store.dispatch(
      returnGroupStudyBooking(returndata, successAlert, errorAlert)
    );
  };

  if (data.groupBookingId) {
    const MySwal = withReactContent(Swal);
    MySwal.fire({
      title: !userData.LastGroupStudyBooking.isConfirmed
        ? data.t("Do you want to confirm your group booking?")
        : data.t("You have a running group study booking."),
      showConfirmButton: false,
      width: 900,
      timer: timers.popup * 1000,
      imageUrl: !userData.LastGroupStudyBooking.isConfirmed ?
        "/images/question.png" : "/images/info.png",
      imageWidth: 110,
      allowOutsideClick: false,
      html: (
        <div className="button-container two">
          <button className="btn btn-primary swal-btn" id="btn3" onClick={() => returnBooking(data, Swal)}>
            {data.t("RETURN")}
          </button>
          <button className="btn btn-primary swal-btn" id="btn4" onClick={() => myBookingModal(data, Swal, "GROUP")}>
            {data.t("BOOKINGS")}
          </button>
          {!userData.LastGroupStudyBooking.isConfirmed && (
            <>
              <button className="btn btn-primary swal-btn" id="btn1" onClick={() => groupBookingConfirmModal(data)}>
                {data.t('YES')}
              </button>

            </>
          )}
          <button className="btn btn-primary swal-btn format-two" id="btn2" onClick={() => askCarrolBookingConfirm(data, Swal)}>
            {data.t(!userData.LastGroupStudyBooking.isConfirmed ? 'NO' : 'CONTINUE')}
          </button>
        </div>
      ),
    }).then((result) => {
      if (result.dismiss === Swal.DismissReason.timer) {
        logout(data);
      }
    });
  } else {
    askCarrolBookingConfirm(data);
  }
};

const askCarrolBookingConfirm = (data, exSwal) => {
  const timers = store.getState().setupReducer.timers

  if (exSwal) {
    exSwal.close();
  }

  const MySwal = withReactContent(Swal);
  const successAlert = (message) => {
    MySwal.fire({
      title: data.t("Success"),
      html: data.t(message),
      showConfirmButton: false,
      width: 600,
      imageUrl: "/images/check.png",
      timer: timers.success * 1000,
      imageWidth: 110,
      allowOutsideClick: false,
    }).then((result) => {
      if (result.dismiss === Swal.DismissReason.timer) {
        logout(data);
      }
    });
  }

  const errorAlert = (message) => {
    MySwal.fire({
      title: data.t('Opps!'),
      showConfirmButton: false,
      width: 600,
      imageUrl: "/images/cancel.png",
      timer: timers.error * 1000,
      imageWidth: 110,
      html: <div>{data.t(message)}</div>,
      allowOutsideClick: false,
    }).then((result) => {
      if (result.dismiss === Swal.DismissReason.timer) {
        logout(data);
      }
    });
  }

  const confirmBooking = (data, exSwal) => {
    if (exSwal) {
      exSwal.close();
    }
    const cunfirmationData = {
      bookingId: data.carrolBookingId,
      userId: data.userId
    }
    store.dispatch(cunfirmCarrelBooking(cunfirmationData, successAlert, errorAlert))
  }

  const returnBooking = (data, exSwal) => {
    if (exSwal) {
      exSwal.close()
    }
    const returnData = {
      bookingId: data.carrolBookingId,
      userId: data.userId
    }
    store.dispatch(returnCarrelBooking(returnData, successAlert, errorAlert))
  }

  if (data.carrolBookingId) {
    const MySwal = withReactContent(Swal);
    MySwal.fire({
      title: data.carrelConfirmation ? data.t("Do you want to confirm your carrel booking?") :
        data.t("You have a running carrol booking.")
      ,
      showConfirmButton: false,
      width: 900,
      timer: timers.popup * 1000,
      allowOutsideClick: false,
      imageUrl: data.carrelConfirmation ? "/images/question.png" : "/images/info.png",
      imageWidth: 110,
      html: (
        <div className="button-container two">
          <button className="btn btn-primary swal-btn" id="btn3" onClick={() => returnBooking(data, Swal)}>
            {data.t("RETURN")}
          </button>
          <button className="btn btn-primary swal-btn" id="btn4" onClick={() => myBookingModal(data, Swal, 'CARREL')}>
            {data.t("BOOKINGS")}
          </button>
          {
            data.carrelConfirmation && (
              <button className="btn btn-primary swal-btn" id="btn1" onClick={() => confirmBooking(data, Swal)} >
                {data.t('YES')}
              </button>
            )
          }
          <button className="btn btn-primary swal-btn format-two" id="btn2" onClick={() => askSeatBookingConfirm(data, Swal)}>
            {data.t(data.carrelConfirmation ? 'NO' : 'CONTINUE')}
          </button>
        </div>
      ),
    }).then((result) => {
      if (result.dismiss === Swal.DismissReason.timer) {
        logout(data);
      }
    });
  } else {
    askSeatBookingConfirm(data);
  }
};

const askSeatBookingConfirm = (data, exSwal) => {
  const timers = store.getState().setupReducer.timers;
  const userData = store.getState().authReducer.userData;
  if (exSwal) {
    exSwal.close();
  }

  const seatChange = (data, exSwal) => {
    exSwal.close();
    data.navigate(
      `/floor/${data.floorID}?type=change&&zoneId=${data.zoneId}&&roomId=${data.roomId}`
    );
  };

  const returnBooking = (data, exSwal) => {
    const MySwal = withReactContent(Swal);

    if (!data.kioskConfig.allowReturn) {
      return MySwal.fire({
        title: data.t('Opps!'),
        showConfirmButton: false,
        width: 600,
        imageUrl: "/images/cancel.png",
        timer: timers.error * 1000,
        imageWidth: 110,
        html: <div>{data.t('Seat return is not allowed!')}</div>,
      }).then((result) => {
        if (result.dismiss === Swal.DismissReason.timer) {
          logout(data);
        }
      })
    }

    const successAlert = () => {
      MySwal.fire({
        title: data.t("Success"),
        html: <div>{data.t("Seat booking returned successfully")}</div>,
        showConfirmButton: false,
        width: 600,
        imageUrl: "/images/check.png",
        timer: timers.success * 1000,
        imageWidth: 110,
        allowOutsideClick: false,
      }).then((result) => {
        if (result.dismiss === Swal.DismissReason.timer) {
          logout(data);
        }
      })
    }

    const errorAlert = (message) => {
      MySwal.fire({
        title: data.t('Opps!'),
        showConfirmButton: false,
        width: 600,
        imageUrl: "/images/cancel.png",
        timer: timers.error * 1000,
        imageWidth: 110,
        html: <div>{data.t(message)}</div>,
        allowOutsideClick: false,
      }).then((result) => {
        if (result.dismiss === Swal.DismissReason.timer) {
          logout(data);
        }
      });
    };

    const returnData = {
      bookingId: data.ReadingRoomBookingId,
    };
    store.dispatch(returnSeatBooking(returnData, successAlert, errorAlert));
    logout(data, Swal);
  };

  const extendBooking = (data, exSwal) => {
    const MySwal = withReactContent(Swal);

    if (!data.kioskConfig.allowExtension) {
      return MySwal.fire({
        title: data.t('Opps!'),
        showConfirmButton: false,
        width: 600,
        imageUrl: "/images/cancel.png",
        timer: timers.error * 1000,
        imageWidth: 110,
        html: <div>{data.t('Seat extension is not allowed!')}</div>,
        allowOutsideClick: false,
      }).then((result) => {
        if (result.dismiss === Swal.DismissReason.timer) {
          logout(data);
        }
      })
    }

    const successAlert = (bookingData) => {
      // MySwal.fire({
      //   title: data.t("Success"),
      //   html: data.t("Seat booking extended successfully"),
      //   showConfirmButton: false,
      //   width: 600,
      //   imageUrl: "/images/check.png",
      //   timer: timers.success * 1000,
      //   imageWidth: 110,
      //   allowOutsideClick: false,
      // }).then((result) => {
      //   if (result.dismiss === Swal.DismissReason.timer) {
      //     logout(data);
      //   }
      // })

      const dispatch = (func) => store.dispatch(func)

      SeatBookingModel(data.navigate,
        bookingData,
        userData,
        data.t,
        timers,
        dispatch,
        store.getState().setupReducer.selectedLanguage)

    };

    const errorAlert = (message) => {
      MySwal.fire({
        title: data.t('Opps!'),
        showConfirmButton: false,
        width: 600,
        imageUrl: "/images/cancel.png",
        timer: timers.error * 1000,
        imageWidth: 110,
        html: <div>{data.t(message)}</div>,
        allowOutsideClick: false,
      }).then((result) => {
        if (result.dismiss === Swal.DismissReason.timer) {
          logout(data);
        }
      })
    };
    const extensionData = {
      bookingId: data.ReadingRoomBookingId,
    }
    store.dispatch(extendSeatBooking(extensionData, successAlert, errorAlert));
  }

  if (data.ReadingRoomBookingId) {
    const MySwal = withReactContent(Swal);
    MySwal.fire({
      imageUrl: '/images/info.png',
      imageAlt: 'Info',
      title: data.t("You have one running booking"),
      showConfirmButton: false,
      width: 1000,
      imageWidth: 120,
      timer: timers.popup * 1000,
      allowOutsideClick: false,
      html: (
        <div className="button-container">
          <button className="btn btn-primary swal-btn" id="btn1" onClick={() => returnBooking(data, Swal)}>
            {data.t("RETURN")}
          </button>
          {data?.allowChangeReadingRoomSeat ? (
            <button className="btn btn-primary  swal-btn" id="btn2" onClick={() => seatChange(data, Swal)}>
              {data.t("CHANGE")}
            </button>
          ) : null}
          {data?.allowExtendReadingRoom ? (
            <button className="btn  btn-primary  swal-btn" id="btn3" onClick={() => extendBooking(data, Swal)}>
              {data.t("EXTEND")}
            </button>
          ) : null}
          <button className="btn btn-primary swal-btn" id="btn4" onClick={() => myBookingModal(data, Swal, "SEAT")}>
            {data.t("BOOKINGS")}
          </button>
          <button className="btn btn-secondary swal-btn format-two" id="btn5" onClick={() => logout(data, Swal)}>
            {data.t("LOGOUT")}
          </button>
        </div>
      ),
    }).then((result) => {
      if (result.dismiss === Swal.DismissReason.timer) {
        logout(data);
      }
    });
  } else {
    data.navigate("/floor/" + data.nextFloorId);
  }
}
