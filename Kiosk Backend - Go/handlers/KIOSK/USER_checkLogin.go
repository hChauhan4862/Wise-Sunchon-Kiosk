package KIOSK_HANDLER

import (
	DB "WISE_SOFTWARE/DB"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func USER_validateLogin(c *gin.Context) {
	type result_USER_validateLogin struct {
		SCHOOL_NO          string `json:"SCHOOL_NO"`
		NAME               string `json:"NAME"`
		USER_POSITION      string `json:"USER_POSITION"`
		USER_POSITION_NAME string `json:"USER_POSITION_NM"`
		DEPT_CODE          string `json:"DEPT_CODE"`
		DEPT_NAME          string `json:"DEPT_NAME"`
		DATE_EXPRD         string `json:"DATE_EXPRD"`
		STATUS             string `json:"STATUS"`
		STATUS_NAME        string `json:"STATUS_NAME"`
		MOBILE             string `json:"MOBILE"`
		EMAIL              string `json:"EMAIL"`
		IS_EXPIRED         int64  `json:"IS_EXPIRED"`
	}

	type viewSeatBooking struct {
		Bseqno          int64      `gorm:"column:bseqno"`
		Seatno          string     `gorm:"column:seatno"`
		Usestart        time.Time  `gorm:"column:usestart"`
		Useexpire       time.Time  `gorm:"column:useexpire"`
		Status          int64      `gorm:"column:status"`
		StatusName      string     `gorm:"column:STATUS_NAME"`
		StatusNameEn    string     `gorm:"column:STATUS_NAME_EN"`
		FloorName       string     `gorm:"column:FLOOR_NAME"`
		FloorNameEn     string     `gorm:"column:FLOOR_NAME_EN"`
		RoomName        string     `gorm:"column:ROOM_NAME"`
		RoomNameEn      string     `gorm:"column:ROOM_NAME_EN"`
		SectorName      string     `gorm:"column:SECTOR_NAME"`
		SectorNameEn    string     `gorm:"column:SECTOR_NAME_EN"`
		SeatName        string     `gorm:"column:SEAT_NAME"`
		SeatNameEn      string     `gorm:"column:SEAT_NAME_EN"`
		SeatVname       string     `gorm:"column:seat_vname"`
		Sectorno        int64      `gorm:"column:sectorno"`
		Roomno          int64      `gorm:"column:roomno"`
		Typeno          int64      `gorm:"column:typeno"`
		ContMin         int64      `gorm:"column:cont_min"`
		MEMBERS         string     `gorm:"column:MEMBERS"`
		RoomCloseTime   time.Time  `gorm:"column:ROOM_CLOSE_TIME"`
		SEATNEXTBOOKING *time.Time `gorm:"column:SEATNEXTBOOKING"`
		ALLOW_ExtendYn  string     `gorm:"column:EXTEND_YN"`
		ALLOW_MoveYn    string     `gorm:"column:MOVE_YN"`
		ALLOW_ReturnYN  string     `gorm:"column:RETURN_YN"`
		ALLOW_CancelYN  string     `gorm:"column:CANCEL_YN"`
		ALLOW_ConfirmYN string     `gorm:"column:CONFIRM_YN"`
		ExtendTill      *time.Time `gorm:"column:EXTEND_TILL"`
	}
	_ = c.MustGet("LIB_NO")

	var schoolID string = c.Params.ByName("userID")

	if schoolID == "=" {
		schoolID = "NEOSCO3"
	}

	var user result_USER_validateLogin
	DB.Conn.Raw(`SELECT alt_pid  SCHOOL_NO,
	username NAME,
	pat_type USER_POSITION,
	pat_name USER_POSITION_NAME,
	dept_code,
	dept_name,
	date_exprd,
	''       STATUS,
	''       STATUS_NAME,
	phone3   MOBILE,
	email,
	CASE WHEN date_exprd < GETDATE() THEN 1 ELSE 0 END IS_EXPIRED
FROM   patmast
WHERE  alt_pid = ?;`, schoolID).Scan(&user)

	if user.SCHOOL_NO == "" {
		c.JSON(http.StatusOK, gin.H{
			"error":      true,
			"error_code": "UL001",
			"result":     "User Not Found",
		})
		return
	}
	if user.IS_EXPIRED == 1 {
		c.JSON(http.StatusOK, gin.H{
			"error":      true,
			"error_code": "UL002",
			"result":     "User Expired",
		})
		return
	}

	isBlocked := 0
	DB.Conn.Raw(`SELECT Count(*)
	FROM   blacklist2
	WHERE  schoolno = ?
		   AND status > 0
		   AND blockstart <= CONVERT(VARCHAR, Dateadd(minute, 30, Getdate())
		   AND blockexpire >= CONVERT(VARCHAR, Dateadd(minute, 30, Getdate());`, schoolID).Scan(&isBlocked)

	if isBlocked > 0 {
		c.JSON(http.StatusOK, gin.H{
			"error":      true,
			"error_code": "UL003",
			"result":     "User Blocked",
		})
		return
	}

	//////////////////////////         CHECK BOOKING INFO      ////////////////////////////

	var booking []viewSeatBooking
	DB.Conn.Raw(`SELECT *,
	CASE
		WHEN EXTEND_YN = 'N' THEN NULL
		ELSE 
			CASE
				WHEN DATEADD(MINUTE, cont_min, GETDATE()) > ROOM_CLOSE_TIME THEN 
					CASE
						WHEN SEATNEXTBOOKING IS NOT NULL AND ROOM_CLOSE_TIME > SEATNEXTBOOKING THEN SEATNEXTBOOKING
						ELSE ROOM_CLOSE_TIME
					END
				ELSE 
					CASE
						WHEN SEATNEXTBOOKING IS NOT NULL AND DATEADD(MINUTE, cont_min, GETDATE()) > SEATNEXTBOOKING THEN SEATNEXTBOOKING
						ELSE DATEADD(MINUTE, cont_min, GETDATE())
					END
			END
		END AS EXTEND_TILL
FROM (
 SELECT bseqno,
       seatno,
       usestart,
       useexpire,
       status,
       status_name STATUS_NAME,
       status_en_name STATUS_NAME_EN,
       floor_name FLOOR_NAME,
       floor_en_name FLOOR_NAME_EN,
       room_name ROOM_NAME,
       room_en_name ROOM_NAME_EN,
       sector_name SECTOR_NAME,
       sector_en_name SECTOR_NAME_EN,
       seat_name SEAT_NAME,
       seat_en_name SEAT_NAME_EN,
       seat_vname,
       sectorno,
       roomno,
       typeno,
       cont_min,
	   MEMBERS,
       (SELECT closetime
        FROM   dbo.Getusetimeymd(CONVERT(VARCHAR, V.useexpire, 112), V.roomno))
       AS
       ROOM_CLOSE_TIME,
       (SELECT TOP 1 usestart
        FROM   view_seat_booking
        WHERE  seatno = V.seatno
               AND usestart >= V.useexpire
               AND status IN ( 2, 3 )
        ORDER  BY usestart)
       SEATNEXTBOOKING,
       CASE
         WHEN status = 3
              AND typeno NOT IN ( 10, 22 )
              AND can_cont_min > 0
              AND Datediff(n, Getdate(), useexpire) <= can_cont_min
              AND Datediff(n, Getdate(), useexpire) >= 0
              AND extend_cnt < max_cont_cnt THEN 'Y'
         ELSE 'N'
       END
       AS EXTEND_YN,
       CASE
         WHEN status = 3 THEN 'Y'
         ELSE 'N'
       END
       AS MOVE_YN,
       CASE
         WHEN status = 3 THEN 'Y'
         ELSE 'N'
       END
       AS RETURN_YN,
	   CASE
         WHEN status = 2 THEN 'Y'
         ELSE 'N'
       END
       AS CANCEL_YN,
	   CASE
	   	  WHEN status = 2 AND (Datediff(n, Getdate(), usestart) <= 30 AND Datediff(n, Getdate(), usestart) > -30) THEN 'Y'
         ELSE 'N'
       END
       AS CONFIRM_YN
FROM   view_seat_booking V
WHERE  ( schoolno IN (SELECT alt_pid
                      FROM   patmast2
                      WHERE  rpst_pers_no IN (SELECT rpst_pers_no
                                              FROM   patmast2
                                              WHERE  alt_pid = ?)
                      GROUP  BY alt_pid)
          OR EXISTS (SELECT TOP 1 'EXIST'
                     FROM   seat_grbooking_id
                     WHERE  bseqno = V.bseqno
                            AND schoolno = ?) )
       AND status IN ( 2, 3 ) AND TYPENO NOT IN ( 23 )
) AS SubQuery
ORDER BY usestart ASC;`, schoolID, schoolID).Scan(&booking)

	c.JSON(http.StatusOK, gin.H{
		"error":      false,
		"error_code": 200,
		"result":     gin.H{"user": user, "booking": booking},
	})
}
