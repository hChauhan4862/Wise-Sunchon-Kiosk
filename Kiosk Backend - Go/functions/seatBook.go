package function

import (
	DB "WISE_SOFTWARE/DB"
	DBmdl "WISE_SOFTWARE/DB/model"

	// "encoding/json"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Mdl_members struct {
	Schoolno string `json:"schoolno"`
	Name     string `json:"name"`
}

type Mdl_SeatBook struct {
	SeatNo        int64         `json:"seatno"`
	Schoolno      string        `json:"schoolno"`
	TimeStart     time.Time     `json:"timestart"`
	TimeEnd       time.Time     `json:"timeend"`
	ISSUE_TYPE_NO int64         `json:"issuetypeno"` // 1 = kiosk, 2 = pc, 3 = mobile
	KioskNo       int64         `json:"kioskno"`     // optional
	IssueFrom     int64         `json:"issuefrom"`   // optional
	IssueDetail   string        `json:"issuedetail"` // optional
	Purpose       string        `json:"purpose"`     // optional
	Tel           string        `json:"tel"`         // optional
	Email         string        `json:"email"`       // optional
	Members       []Mdl_members `json:"members"`     // dict of members (optional)
}

func SeatBook(d Mdl_SeatBook) (string, error) {
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

	type SeatKiosk struct {
		Kioskno       int64  `json:"kioskno"`
		Issuefrom     int64  `json:"issuefrom"`
		Issuedetail   string `json:"issuedetail"`
		LibNo         int64  `json:"LibNo"`
		AssignLibOnly int64  `json:"AssignLibOnly"`
		Assignable    int64  `json:"Assignable"`
	}
	d.TimeEnd = d.TimeEnd.Add(-1 * time.Second) // add duration to start time
	/* -------------------------------------- Validate Data and Get Lock  --------------------------------------*/
	LOCKS := []string{}
	LOCKS = append(LOCKS, fmt.Sprintf("SEAT#%v", d.SeatNo))
	LOCKS = append(LOCKS, fmt.Sprintf("USER#%v", d.Schoolno))
	for _, member := range d.Members {
		LOCKS = append(LOCKS, fmt.Sprintf("USER#%v", member.Schoolno))
	}
	if ok := LockGet(LOCKS); !ok {
		return "LCK500", fmt.Errorf("please try in few seconds")
	}
	defer LockRelease(LOCKS)

	/*-------------------------------------- Validate Time  --------------------------------------*/

	bookTimeStart := d.TimeStart // BOOKING TIME START IN LOCAL TIME
	bookTimeEnd := d.TimeEnd     // BOOKING TIME END IN LOCAL TIME
	bookTimeCurr := time.Now()   // BOOKING TIME CURRENT IN LOCAL TIME

	bookTimeStartSTR := bookTimeStart.Format("2006-01-02 15:04:05") // BOOKING TIME START IN STRING FORMAT
	bookTimeEndSTR := bookTimeEnd.Format("2006-01-02 15:04:05")     // BOOKING TIME END IN STRING FORMAT

	// bookTimeStartUTC := bookTimeStart.UTC() // BOOKING TIME START IN UTC TIME
	// bookTimeEndUTC := bookTimeEnd.UTC()     // BOOKING TIME END IN UTC TIME
	bookTimeStartUTC, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", bookTimeStartSTR+" +0000 UTC") // BOOKING TIME START IN UTC TIME
	bookTimeEndUTC, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", bookTimeEndSTR+" +0000 UTC")     // BOOKING TIME END IN UTC TIME

	bookTimeDate := bookTimeStart.Format("20060102")      // BOOKING TIME DATE IN STRING FORMAT
	bookTimeStartHHMMSS := bookTimeStart.Format("150405") // BOOKING TIME START IN STRING FORMAT
	bookTimeEndHHMMSS := bookTimeEnd.Format("150405")     // BOOKING TIME END IN STRING FORMAT

	bookTimeDur := bookTimeEnd.Sub(bookTimeStart) // get duration in nanoseconds

	fmt.Println("BOOKING TIME START   :: ", bookTimeStart)
	fmt.Println("BOOKING TIME END     :: ", bookTimeEnd)
	fmt.Println("BOOKING TIME CURRENT :: ", bookTimeCurr)

	fmt.Println("BOOKING TIME START STR   :: ", bookTimeStartSTR)
	fmt.Println("BOOKING TIME END STR     :: ", bookTimeEndSTR)

	fmt.Println("BOOKING TIME START UTC   :: ", bookTimeStartUTC)
	fmt.Println("BOOKING TIME END UTC     :: ", bookTimeEndUTC)

	fmt.Println("BOOKING TIME DATE         :: ", bookTimeDate)
	fmt.Println("BOOKING TIME START HHMMSS :: ", bookTimeStartHHMMSS)
	fmt.Println("BOOKING TIME END HHMMSS   :: ", bookTimeEndHHMMSS)

	fmt.Println("BOOKING TIME DURATION     :: ", bookTimeDur)

	if bookTimeStart.After(bookTimeEnd) || time.Now().After(bookTimeStart.Add(2*time.Minute)) {
		return "SB001", fmt.Errorf("incorrect time formats")
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
	WHERE  school_no = ? AND date_exprd >= CONVERT(varchar, getdate(), 112) AND STATUS='Y' AND STATUS_NAME='Y'`, d.Schoolno).Scan(&seatUserdata)

	if seatUserdata.Status != "Y" {
		return "SB002", fmt.Errorf("user not found or Expired")
	}

	/*-------------------------------------- Check if User Blocked  --------------------------------------*/

	isBlocked := 0
	DB.Conn.Raw(`SELECT Count(*)
	FROM   blacklist2
	WHERE  schoolno = ?
		   AND status > 0
		   AND blockstart <= CONVERT(VARCHAR, Dateadd(minute, 30, Getdate())
		   AND blockexpire >= CONVERT(VARCHAR, Dateadd(minute, 30, Getdate());`, d.Schoolno).Scan(&isBlocked)

	if isBlocked > 0 {
		return "SB003", fmt.Errorf("user is blocked")
	}

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
	WHERE  seatno = ?;`, bookTimeDate, bookTimeDate, d.SeatNo).Scan(&seat)

	// temp, _ := json.Marshal(seat)
	// fmt.Println("SEAT DETAILS:: ", string(temp))

	if seat.Seatno == 0 {
		return "SB004", fmt.Errorf("seat not found")
	}

	if seat.Status != 1 && seat.Status != 2 {
		return "SB005", fmt.Errorf("seat is not available")
	}

	if bookTimeStartUTC.Before(seat.Opentime) || bookTimeEndUTC.After(seat.Closetime) {
		return "SB006", fmt.Errorf("invalid operation hours")
	}

	if bookTimeDur > time.Duration(seat.UseMin)*time.Minute {
		return "SB007", fmt.Errorf("maximum usage time exceeded")
	}

	if d.ISSUE_TYPE_NO == 1 && seat.ASSIGN_YN != "Y" {
		return "SB008", fmt.Errorf("seat is not available for kiosk booking")
	}

	if d.ISSUE_TYPE_NO == 3 && seat.MOBILE_ASSIGN_YN != "Y" && seat.MOBILE_BOOKING_YN != "Y" {
		return "SB009", fmt.Errorf("seat is not available for mobile booking")
	}

	if len(d.Members) > seat.GrMax || len(d.Members) < seat.GrMin {
		return "SB010", fmt.Errorf("invalid number of members")
	}

	/*-------------------------------------- USER BOOKINGS IN THIS TIME FRAME  --------------------------------------*/

	typeNotIn := "NOT IN (22, 10)"
	if seat.Typeno == 10 {
		typeNotIn = "IN (10)"
	} else if seat.Typeno == 22 {
		typeNotIn = "IN (22)"
	}

	var c int64
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
				AND useexpire >= ? )
			OR ( usestart <= ?
					AND useexpire >= ? )
			OR ( usestart >= ?
					AND useexpire <= ? ) )
		AND iscanceled = 0
		AND typeno `+typeNotIn+` `, d.Schoolno, d.Schoolno, bookTimeStartSTR, bookTimeStartSTR, bookTimeEndSTR, bookTimeEndSTR, bookTimeStartSTR, bookTimeEndSTR).Scan(&c)

	if c > 0 {
		return "SB011", fmt.Errorf("user has another booking for this specific time")
	}

	/*-------------------------------------- SEAT BOOKING HISTORY IN THIS TIME FRAME  --------------------------------------*/

	var seatBookingCount int64
	DB.Conn.Raw(`
	SELECT count(BSEQNO) booking_cnt
		FROM   seat_booking
		WHERE  seatno = ?
       		AND ( ( usestart <= ?
               		AND useexpire >= ? )
              	OR ( usestart <= ?
                   	AND useexpire >= ? )
              	OR ( usestart >= ?
                   AND useexpire <= ? ) )
       	AND iscanceled = 0
       AND status IN ( 1, 2, 3 )`, d.SeatNo, bookTimeStartSTR, bookTimeStartSTR, bookTimeEndSTR, bookTimeEndSTR, bookTimeStartSTR, bookTimeEndSTR).Row().Scan(&seatBookingCount)

	if seatBookingCount > 0 {
		return "SB012", fmt.Errorf("seat is already booked for this time")
	}

	/*-------------------------------------- CHECK USER POSITION ELIGIBILITY  --------------------------------------*/
	var seatAllowedPosition string
	DB.Conn.Raw(`SELECT user_pos FROM   seat_room2 WHERE  roomno = ?`, seat.Roomno).Scan(&seatAllowedPosition)

	fmt.Println("SEAT ALLOWED POSITION:: ", seatAllowedPosition)

	if !strings.Contains(seatAllowedPosition, "(*)") && !strings.Contains(seatAllowedPosition, seatUserdata.UserPosition) {
		return "SB013", fmt.Errorf("user position is not allowed to book this seat")
	}

	/* ------------------------------------ CHECK IF THIS IS LIBRARY CLOSE DAY  ----------------------------------- */
	var closeTitle string = "NA"
	DB.Conn.Raw(`SELECT title title
	FROM   lib_close_day
	WHERE  ( 
			    ( close_dt_fr <= ? AND close_tm_fr <= ? ) 
			AND ( close_dt_to >= ? AND close_tm_to > ? )
			 OR ( close_dt_fr <= ? AND close_tm_fr < ? )
		    AND ( close_dt_to >= ? AND close_tm_to >= ? ) 
	) AND ( libno = ? OR libno = 0 )
	  AND ( floor = ? OR floor = 0 )
	  AND ( roomno = ? OR roomno = 0 )`,
		bookTimeDate,
		bookTimeStartHHMMSS,
		bookTimeDate,
		bookTimeStartHHMMSS,
		bookTimeDate,
		bookTimeEndHHMMSS,
		bookTimeDate,
		bookTimeEndHHMMSS,
		seat.Libno, seat.FLOORNO, seat.Roomno).Scan(&closeTitle)

	fmt.Println("CLOSE TITLE:: ", closeTitle)

	if closeTitle != "NA" {
		return "SB014", fmt.Errorf("library is closed on this day")
	}

	/* ------------------------------------ GET KIOSK INFORMATION  ----------------------------------- */
	if d.ISSUE_TYPE_NO == 1 {
		var seatKiosk SeatKiosk
		DB.Conn.Raw(`SELECT kioskno,
			issuefrom,
			issuedetail,
			LibNo,
			AssignLibOnly,
			Assignable
		FROM   seat_kiosk2
		WHERE  kioskno = ?`, d.KioskNo).Scan(&seatKiosk)

		// temp, _ = json.Marshal(seatKiosk)
		// fmt.Println("SEAT DETAILS:: ", string(temp))

		if seatKiosk.Kioskno == 0 {
			return "SB015", fmt.Errorf("kiosk not found")
		}

		if seatKiosk.Assignable == 0 {
			return "SB016", fmt.Errorf("kiosk is not available")
		}

		if seatKiosk.AssignLibOnly == 1 && seatKiosk.LibNo != seat.Libno {
			return "SB017", fmt.Errorf("kiosk is not available for this seat")
		}

		d.IssueFrom = seatKiosk.Issuefrom
		d.IssueDetail = seatKiosk.Issuedetail
	}

	/*-------------------------------------- BOOK SEAT  --------------------------------------*/
	var bseqno int64 = -1
	var status int64 = 3
	var expirereason int64 = 1
	var seatUseStart time.Time = bookTimeStart
	DB.Conn.Raw(`SELECT Max(bseqno) + 1 maxno FROM   seat_booking`).Scan(&bseqno)

	if bseqno == -1 {
		return "SB018", fmt.Errorf("error generating booking number")
	}

	if seat.Typeno == 10 || seat.Typeno == 22 {
		expirereason = 0
		status = 2
		seatUseStart = time.Time{}
	}

	/*-------------------------------------- BOOK SEAT  --------------------------------------*/

	// err := DB.Conn.Transaction(func(tx *gorm.DB) error {
	err := DB.Conn.Session(&gorm.Session{DryRun: false}).Transaction(func(tx *gorm.DB) error {
		var zero int64 = 0
		var now time.Time = time.Now()
		if err := tx.Create(&DBmdl.SeatBooking{
			BSEQNO:                 bseqno,
			STATUS:                 &status,
			SCHOOLNO:               &d.Schoolno,
			SEATNO:                 &d.SeatNo,
			ISCANCELED:             &zero,
			ISSUEFROM:              &d.IssueFrom,
			ISSUEDETAIL:            &d.IssueDetail,
			USESTART:               &d.TimeStart,
			USEEXPIRE:              &d.TimeEnd,
			EXPIREREASON:           &expirereason,
			QUOTAMIN:               &zero,
			PRINTCNT:               &zero,
			REGTIME:                &now,
			EXTENDMIN:              &zero,
			EXTENDCNT:              &zero,
			STARTTIME:              &seatUseStart,
			USERNAME:               &seatUserdata.Name,
			USERPOSITION:           &seatUserdata.UserPosition,
			COMPANYCODE:            &seatUserdata.CompanyCode,
			DEPTCODE:               &seatUserdata.DeptCode,
			ISADMINBOOKING:         &zero,
			CHECKGATEIN:            &zero,
			PCLOGINCHECKSTATUS:     &zero,
			EXTENDMSGSTATUS:        &zero,
			ISSUETYPENO:            &d.ISSUE_TYPE_NO,
			BEFORESEATRETURNSTATUS: &zero,
		}).Error; err != nil {
			return err
		}

		var logstatecode int64 = 2
		if seat.Typeno == 10 || seat.Typeno == 22 {
			logstatecode = 1
		}

		if err := tx.Create(&DBmdl.SeatBookingLog{
			BSEQNO:      &bseqno,
			SCHOOLNO:    &d.Schoolno,
			SEATNO:      &d.SeatNo,
			STATUS:      &logstatecode,
			USESTART:    &d.TimeStart,
			USEEXPIRE:   &d.TimeEnd,
			STARTTIME:   &seatUseStart,
			NEWSEATNO:   &zero,
			EXTENDMIN:   &zero,
			EXTENDCNT:   &zero,
			LOGTIME:     &now,
			ISSUEFROM:   &d.IssueFrom,
			ISSUEDETAIL: &d.IssueDetail,
			ISSUETYPENO: &d.ISSUE_TYPE_NO,
		}).Error; err != nil {
			return err
		}

		var minutes time.Duration = 0
		tx.Raw(`select TIMEOUT_GATEIN_USER_MIN from LIBCONFIG where STATUS = 0`, seat.Roomno).Row().Scan(&minutes)
		var ns time.Duration = (minutes * time.Minute) + (2 * time.Second)

		// Add Backgroud Job
		if status == 2 {
			TaskAlarm(bookTimeEnd.Add(ns), "BOOKING#RETURN#%v#%v", bseqno, 3)   // RETURN AFTER END TIME, MANUAL CONFIRM
			TaskAlarm(bookTimeStart.Add(ns), "BOOKING#RETURN#%v#%v", bseqno, 1) // RETURN AFTER 30 MINUTES OF NOT CONFIRMING
		} else {
			TaskAlarm(bookTimeEnd.Add(ns), "BOOKING#RETURN#%v#%v", bseqno, 2) // RETURN AFTER END TIME AUTO CONFIRM
		}
		return nil
	})

	if err != nil {
		return "ERR99", fmt.Errorf("something went wrong")
	}

	return fmt.Sprint(bseqno), nil

}
