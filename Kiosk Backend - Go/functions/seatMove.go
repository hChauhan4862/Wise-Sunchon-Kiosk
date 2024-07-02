package function

import (
	DB "WISE_SOFTWARE/DB"
	DBmdl "WISE_SOFTWARE/DB/model"
	"encoding/json"

	// "encoding/json"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Mdl_SeatMove struct {
	Bseqno        int64  `json:"bseqno"`
	New_seatNo    int64  `json:"new_seatNo"`
	ISSUE_TYPE_NO int64  `json:"issuetypeno"` // 1 = kiosk, 2 = pc, 3 = mobile
	KioskNo       int64  `json:"kioskno"`     // optional
	IssueFrom     int64  `json:"issuefrom"`   // optional
	IssueDetail   string `json:"issuedetail"` // optional
}

func SeatMove(d Mdl_SeatMove) (string, error) {
	type currentSeatBooking struct {
		BSeqNo       int64     `gorm:"column:bseqno"`
		SeatNo       int64     `gorm:"column:seatno"`
		UserName     string    `gorm:"column:user_name"`
		Usestart     time.Time `gorm:"column:usestart"`
		Useexpire    time.Time `gorm:"column:useexpire"`
		FloorName    string    `gorm:"column:FLOOR_NAME"`
		RoomName     string    `gorm:"column:ROOM_NAME"`
		SectorName   string    `gorm:"column:SECTOR_NAME"`
		SeatName     string    `gorm:"column:SEAT_NAME"`
		SeatVname    string    `gorm:"column:seat_vname"`
		Status       int       `gorm:"column:status"`
		Schoolno     string    `gorm:"column:schoolno"`
		Starttime    time.Time `gorm:"column:starttime"`
		Issuefrom    int64     `gorm:"column:issuefrom"`
		Issuedetail  string    `gorm:"column:issuedetail"`
		IssueTypeNo  int       `gorm:"column:issue_type_no"`
		Printcnt     int64     `gorm:"column:printcnt"`
		ExtendMin    int64     `gorm:"column:extend_min"`
		ExtendCnt    int64     `gorm:"column:extend_cnt"`
		CanContMin   int64     `gorm:"column:can_cont_min"`
		ContMin      int64     `gorm:"column:cont_min"`
		MaxContCount int64     `gorm:"column:max_cont_cnt"`
		TypeNo       int64     `gorm:"column:typeno"`
		GrMin        int64     `gorm:"column:gr_min"`
		GrMax        int64     `gorm:"column:gr_max"`
		UseMin       int64     `gorm:"column:use_min"`
	}
	type mdl_SeatUserdata3 struct {
		UserID       int    `gorm:"column:pid;userid"`
		SchoolNo     string `gorm:"column:school_no;schoolno"`
		Name         string
		UserPosition string `gorm:"column:pat_type;user_position"`
		DeptCode     string `gorm:"column:dept_code"`
		CompanyCode  string `gorm:"column:company_code"`
		DateExprd    string `gorm:"column:date_exprd;date_exprd"`
		Status       string
		StatusName   string `gorm:"column:status_name"`
	}
	type mdl_SeatSeatView struct {
		Seatno            int64     `json:"seatno"`
		Name              string    `json:"name"`
		Vname             string    `json:"vname"`
		Sectorno          int       `json:"sectorno"`
		BOOKING_YN        string    `json:"booking_yn"`
		ASSIGN_YN         string    `json:"assign_yn"`
		MOBILE_BOOKING_YN string    `json:"mobile_booking_yn"`
		MOBILE_ASSIGN_YN  string    `json:"mobile_assign_yn"`
		Status            int       `json:"status"`
		Libno             int64     `json:"libno"`
		GrMin             int       `json:"gr_min"`
		GrMax             int       `json:"gr_max"`
		UseMin            int       `json:"use_min"`
		Typeno            int       `json:"typeno"`
		TypeName          string    `json:"type_name"`
		Roomno            int       `json:"roomno"`
		FLOORNO           int       `json:"floorno"`
		FloorName         string    `json:"floor_name"`
		RoomName          string    `json:"room_name"`
		SectorName        string    `json:"sector_name"`
		Opentime          time.Time `json:"opentime"`
		Closetime         time.Time `json:"closetime"`
	}

	type seatKiosk struct {
		Kioskno       int64  `json:"kioskno"`
		Issuefrom     int64  `json:"issuefrom"`
		Issuedetail   string `json:"issuedetail"`
		LibNo         int64  `json:"LibNo"`
		AssignLibOnly int64  `json:"AssignLibOnly"`
		Assignable    int64  `json:"Assignable"`
	}
	/* -------------------------------------- Get Lock  --------------------------------------*/
	if ok := LockGet([]string{fmt.Sprintf("BSEQNO#%v", d.Bseqno)}); !ok {
		return "LCK500", fmt.Errorf("please try in few seconds")
	}
	defer LockRelease([]string{fmt.Sprintf("BSEQNO#%v", d.Bseqno)})

	/*-------------------------------------- Check if Booking is Valid  --------------------------------------*/
	var CB currentSeatBooking
	DB.Conn.Raw(`SELECT bseqno,
	seatno,
	user_name,
	usestart,
	useexpire,
	floor_name                       FLOOR_NAME,
	room_name                        ROOM_NAME,
	sector_name                      SECTOR_NAME,
	seat_name                        SEAT_NAME,
	seat_vname,
	status,
	schoolno,
	starttime,
	issuefrom,
	issuedetail,
	issue_type_no,
	printcnt,
	extend_min,
	extend_cnt,
	can_cont_min,
	cont_min,
	max_cont_cnt,
	typeno,
	gr_min,
	gr_max,
	use_min
FROM   view_seat_booking V
WHERE  bseqno = ? AND status IN ( 3 ) AND iscanceled = 0`, d.Bseqno).Scan(&CB)

	if CB.BSeqNo != d.Bseqno {
		return "SM001", fmt.Errorf("booking not found")
	}

	if CB.SeatNo == d.New_seatNo {
		return "SM015", fmt.Errorf("seat moving not allowed for same seat")
	}

	bookTimeNow := time.Now() // BOOKING TIME CURRENT IN LOCAL TIME
	bookTimeNowSTR := bookTimeNow.Format("2006-01-02 15:04:05")
	bookTimeNowUTC, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", bookTimeNowSTR+" +0000 UTC")

	bookTimeStartUTC := CB.Usestart
	bookTimeStartSTR := bookTimeStartUTC.Format("2006-01-02 15:04:05")

	bookTimeEndUTC := CB.Useexpire
	bookTimeEndSTR := bookTimeEndUTC.Format("2006-01-02 15:04:05")

	bookTimeDate := bookTimeStartUTC.Format("20060102") // BOOKING TIME DATE IN STRING FORMAT
	// bookTimeStartHHMMSS := bookTimeStartUTC.Format("150405") // BOOKING TIME START IN STRING FORMAT
	// bookTimeEndHHMMSS := bookTimeEndUTC.Format("150405")     // BOOKING TIME END IN STRING FORMAT

	if bookTimeEndUTC.Before(bookTimeNowUTC) {
		return "SM002", fmt.Errorf("seat moving not allowed for expired booking")
	}

	/* ------------------------------------ GET LOCK FOR SEAT AND USER  ----------------------------------- */
	if ok := LockGet([]string{fmt.Sprintf("SEAT#%v", d.New_seatNo), fmt.Sprintf("USER#%v", CB.Schoolno)}); !ok {
		return "LCK500", fmt.Errorf("please try in few seconds")
	}
	defer LockRelease([]string{fmt.Sprintf("SEAT#%v", d.New_seatNo), fmt.Sprintf("USER#%v", CB.Schoolno)})

	/*-------------------------------------- Seat Details  --------------------------------------*/

	var seat mdl_SeatSeatView
	DB.Conn.Raw(`SELECT S.seatno,
		S.NAME                                           name,
		S.vname,
		S.sectorno,
		S.BOOKING_YN,
		S.ASSIGN_YN,
		S.MOBILE_BOOKING_YN,
		S.MOBILE_ASSIGN_YN,
		S.status,
		S.libno,
		S.gr_min,
		S.gr_max,
		S.use_min,
		S.typeno,
		S.type_name                                      type_name,
		S.roomno,
		S.floor_name                                     floor_name,
		S.FLOORNO,
		S.room_name                                      room_name,
		S.sector_name                                    sector_name,
		(SELECT opentime
		FROM   dbo.Getusetimeymd(?, S.roomno)) opentime,
		(SELECT closetime
		FROM   dbo.Getusetimeymd(?, S.roomno)) closetime
	FROM   view_seat_seat S
	WHERE  seatno = ?;`, bookTimeDate, bookTimeDate, d.New_seatNo).Scan(&seat)

	temp, _ := json.Marshal(seat)
	fmt.Println("SEAT DETAILS:: ", string(temp))

	if seat.Seatno == 0 {
		return "SM003", fmt.Errorf("seat not found")
	}

	if seat.Status != 1 && seat.Status != 2 {
		return "SM004", fmt.Errorf("seat is not available")
	}

	if bookTimeNowUTC.Before(seat.Opentime) || bookTimeEndUTC.After(seat.Closetime) {
		return "SM005", fmt.Errorf("invalid operation hours")
	}

	// No Need At All Below Commented Code
	bookDur := bookTimeEndUTC.Sub(bookTimeNowUTC) // get duration in nanoseconds
	if bookDur > time.Duration(seat.UseMin)*time.Minute {
		return "SM006", fmt.Errorf("maximum usage time exceeded for selected Seat")
	}

	if d.ISSUE_TYPE_NO == 1 && seat.ASSIGN_YN != "Y" {
		return "SM007", fmt.Errorf("seat is not available for kiosk booking")
	}

	if d.ISSUE_TYPE_NO == 3 && seat.MOBILE_ASSIGN_YN != "Y" && seat.MOBILE_BOOKING_YN != "Y" {
		return "SM008", fmt.Errorf("seat is not available for mobile booking")
	}

	/*-------------------------------------- Check If Seat is Booked by Other  --------------------------------------*/

	var seatBookingCount int
	DB.Conn.Raw(`SELECT Count(*) USECNT
		FROM   view_seat_booking
		WHERE  seatno = ?
		AND ( ( usestart <= ?
			AND useexpire > ? )
		OR ( usestart < ?
				AND useexpire > ? )
		OR ( usestart > ?
				AND useexpire <= ? ) )
	AND iscanceled = 0
	AND status IN ( 2, 3 ) `, d.New_seatNo, bookTimeStartSTR, bookTimeStartSTR, bookTimeEndSTR, bookTimeEndSTR, bookTimeStartSTR, bookTimeEndSTR).Scan(&seatBookingCount)

	if seatBookingCount > 0 {
		return "SM009", fmt.Errorf("seat has been booked by another user")
	}

	/*-------------------------------------- Validate User  --------------------------------------*/
	var seatUserdata mdl_SeatUserdata3
	DB.Conn.Raw(`SELECT pid           userid,
		school_no     schoolno,
		NAME,
		pat_type      user_position,
		dept_code,
		CASE
		WHEN Substring(dept_code, 4, 1) = '_' THEN LEFT(dept_code, 3)
		WHEN Substring(dept_code, 5, 1) = '_' THEN LEFT(dept_code, 4)
		ELSE dept_code
		END           company_code,
		date_exprd    date_exprd,
		status,
		status_name
	FROM   seat_userdata3
	WHERE  school_no = ? AND date_exprd >= CONVERT(varchar, getdate(), 112) AND STATUS='Y' AND STATUS_NAME='Y'`, CB.Schoolno).Scan(&seatUserdata)

	if seatUserdata.Status != "Y" {
		return "SM010", fmt.Errorf("user not found or Expired")
	}

	/*-------------------------------------- Check If User has Booked Another Seat  --------------------------------------*/

	typeNotIn := "NOT IN (22, 10)"
	if CB.TypeNo == 10 {
		typeNotIn = "IN (10)"
	} else if CB.TypeNo == 22 {
		typeNotIn = "IN (22)"
	}

	var userBookingCount int64
	DB.Conn.Raw(`SELECT 
	count(BSEQNO)
		FROM   view_seat_booking V
	WHERE  ( schoolno IN (SELECT DISTINCT p1.alt_pid
		FROM patmast2 p1
		INNER JOIN patmast2 p2 ON p1.rpst_pers_no = p2.rpst_pers_no
		WHERE p2.alt_pid = ?)
		OR ( 0 < (SELECT Count(*)
					FROM   seat_grbooking_id
					WHERE  bseqno = V.bseqno
							AND schoolno = ?) ) )
		AND status IN ( 2, 3 )
		AND ( ( usestart <= ?
				AND useexpire > ? )
			OR ( usestart < ?
					AND useexpire > ? )
			OR ( usestart > ?
					AND useexpire <= ? ) )
		AND iscanceled = 0
		AND typeno `+typeNotIn+` AND BSEQNO != ? `, CB.Schoolno, CB.Schoolno, bookTimeStartSTR, bookTimeStartSTR, bookTimeEndSTR, bookTimeEndSTR, bookTimeStartSTR, bookTimeEndSTR, CB.BSeqNo).Scan(&userBookingCount)

	if userBookingCount > 0 {
		return "SM011", fmt.Errorf("user has another booking at the same time")
	}

	/*-------------------------------------- CHECK USER POSITION ELIGIBILITY  --------------------------------------*/
	var seatAllowedPosition string
	DB.Conn.Raw(`SELECT user_pos FROM   seat_room2 WHERE  roomno = ?`, seat.Roomno).Scan(&seatAllowedPosition)

	fmt.Println("SEAT ALLOWED POSITION:: ", seatAllowedPosition)

	if !strings.Contains(seatAllowedPosition, "(*)") && !strings.Contains(seatAllowedPosition, seatUserdata.UserPosition) {
		return "SM012", fmt.Errorf("user position is not allowed to book this seat")
	}

	/* ------------------------------------ GET KIOSK INFORMATION  ----------------------------------- */
	if d.ISSUE_TYPE_NO == 1 {
		var seatKiosk seatKiosk
		DB.Conn.Raw(`SELECT kioskno,
			issuefrom,
			issuedetail,
			LibNo,
			AssignLibOnly,
			Assignable
		FROM   seat_kiosk2
		WHERE  kioskno = ?`, d.KioskNo).Scan(&seatKiosk)

		if seatKiosk.Kioskno == 0 {
			return "SM013", fmt.Errorf("kiosk not found")
		}

		if seatKiosk.Assignable == 0 {
			return "SM014", fmt.Errorf("kiosk is not available")
		}

		d.IssueFrom = seatKiosk.Issuefrom
		d.IssueDetail = seatKiosk.Issuedetail
	}

	/*-------------------------------------- Extend Booking  --------------------------------------*/

	err := DB.Conn.Session(&gorm.Session{DryRun: false}).Transaction(func(tx *gorm.DB) error {
		var zero int64 = 0
		// var one int64 = 1
		var three int64 = 3

		if err := tx.Model(&DBmdl.SeatBooking{}).Where("bseqno = ?", d.Bseqno).Updates(map[string]interface{}{
			"seatno":   d.New_seatNo,
			"printcnt": gorm.Expr("printcnt + ?", 1),
		}).Error; err != nil {
			return err
		}

		if err := tx.Create(&DBmdl.SeatBookingLog{
			BSEQNO:      &d.Bseqno,
			SCHOOLNO:    &CB.Schoolno,
			SEATNO:      &CB.SeatNo,
			STATUS:      &three,
			USESTART:    &CB.Usestart,
			USEEXPIRE:   &CB.Useexpire,
			STARTTIME:   &CB.Starttime,
			NEWSEATNO:   &d.New_seatNo,
			EXTENDMIN:   &zero,
			EXTENDCNT:   &zero,
			LOGTIME:     &bookTimeNowUTC,
			ISSUEFROM:   &d.IssueFrom,
			ISSUEDETAIL: &d.IssueDetail,
			ISSUETYPENO: &d.ISSUE_TYPE_NO,
		}).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return "ERR99", fmt.Errorf("something went wrong")
	}

	return "SUCCESS", nil
}
