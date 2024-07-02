package function

import (
	DB "WISE_SOFTWARE/DB"
	DBmdl "WISE_SOFTWARE/DB/model"

	// "encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Mdl_SeatExtend struct {
	Bseqno        int64     `json:"bseqno"`
	Minute        int64     `json:"minute"`
	TimeEnd       time.Time `json:"timeend"`
	ISSUE_TYPE_NO int64     `json:"issuetypeno"` // 1 = kiosk, 2 = pc, 3 = mobile
	KioskNo       int64     `json:"kioskno"`     // optional
	IssueFrom     int64     `json:"issuefrom"`   // optional
	IssueDetail   string    `json:"issuedetail"` // optional
}

func SeatExtend(d Mdl_SeatExtend) (string, error) {

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
		CloseTime    time.Time `gorm:"column:CLOSETIME"`
	}

	type seatKiosk struct {
		Kioskno       int64  `json:"kioskno"`
		Issuefrom     int64  `json:"issuefrom"`
		Issuedetail   string `json:"issuedetail"`
		LibNo         int64  `json:"LibNo"`
		AssignLibOnly int64  `json:"AssignLibOnly"`
		Assignable    int64  `json:"Assignable"`
	}

	d.TimeEnd = d.TimeEnd.Add(-1 * time.Second) // add duration to start time
	/* -------------------------------------- Get Lock  --------------------------------------*/
	if ok := LockGet([]string{fmt.Sprintf("BSEQNO#%v", d.Bseqno)}); !ok {
		return "LCK500", fmt.Errorf("please try in few seconds")
	}
	defer LockRelease([]string{fmt.Sprintf("BSEQNO#%v", d.Bseqno)})

	/*-------------------------------------- Validate Time  --------------------------------------*/
	bookTimeEnd := d.TimeEnd                                                                      // BOOKING TIME END IN LOCAL TIME
	bookTimeNow := time.Now()                                                                     // BOOKING TIME CURRENT IN LOCAL TIME
	bookTimeEndSTR := bookTimeEnd.Format("2006-01-02 15:04:05")                                   // BOOKING TIME END IN STRING FORMAT
	bookTimeNowSTR := bookTimeNow.Format("2006-01-02 15:04:05")                                   // BOOKING TIME END IN STRING FORMAT
	bookTimeEndUTC, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", bookTimeEndSTR+" +0000 UTC") // BOOKING TIME END IN UTC TIME
	bookTimeNowUTC, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", bookTimeNowSTR+" +0000 UTC") // BOOKING TIME END IN UTC TIME

	fmt.Println("bookTimeEndSTR:: ", bookTimeEndSTR)

	if bookTimeNow.After(bookTimeEnd) || bookTimeNow.Equal(bookTimeEnd) {
		return "SE001", fmt.Errorf("incorrect time formats")
	}

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
	use_min,
	(SELECT closetime
	 FROM   dbo.Getusetimeymd(CONVERT(VARCHAR, V.useexpire, 112), V.roomno))
									 CLOSETIME
FROM   view_seat_booking V
WHERE  bseqno = ? AND status IN ( 3 ) AND iscanceled = 0`, d.Bseqno).Scan(&CB)

	if CB.BSeqNo != d.Bseqno {
		return "SE002", fmt.Errorf("booking not found")
	}

	bookTimeStartSTR := CB.Usestart.Format("2006-01-02 15:04:05")

	if CB.Useexpire.Before(bookTimeNowUTC) {
		return "SE003", fmt.Errorf("extension not allowed for expired booking")
	}

	if CB.MaxContCount == 0 {
		return "SE004", fmt.Errorf("extension not allowed for this booking")
	}

	if CB.ExtendCnt >= CB.MaxContCount {
		return "SE005", fmt.Errorf("maximum extend count reached")
	}

	if CB.CanContMin < int64(bookTimeNowUTC.Sub(CB.Useexpire).Minutes()) {
		return "SE006", fmt.Errorf("cannot extend yet")
	}

	if d.Minute > CB.ContMin {
		return "SE007", fmt.Errorf("cannot extend for more than certain minutes at a time")
	}

	if CB.Useexpire.After(bookTimeEndUTC) || CB.Useexpire.Equal(bookTimeEndUTC) {
		return "SE008", fmt.Errorf("seat is already booked for this time")
	}

	if CB.CloseTime.Before(bookTimeEndUTC) {
		return "SE009", fmt.Errorf("cannot extend beyond closing time")
	}

	// temp, _ := json.Marshal(CB)
	// fmt.Println("SEAT DETAILS:: ", string(temp))

	/*-------------------------------------- Get Lock for Seat and User  --------------------------------------*/
	if ok := LockGet([]string{fmt.Sprintf("SEAT#%v", CB.SeatNo), fmt.Sprintf("USER#%v", CB.Schoolno)}); !ok {
		return "LCK500", fmt.Errorf("please try in few seconds")
	}
	defer LockRelease([]string{fmt.Sprintf("SEAT#%v", CB.SeatNo), fmt.Sprintf("USER#%v", CB.Schoolno)})

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
	AND status IN ( 2, 3 ) AND BSEQNO != ? `, CB.SeatNo, bookTimeStartSTR, bookTimeStartSTR, bookTimeEndSTR, bookTimeEndSTR, bookTimeStartSTR, bookTimeEndSTR, d.Bseqno).Scan(&seatBookingCount)

	if seatBookingCount > 0 {
		return "SE010", fmt.Errorf("seat has been booked by another user")
	}

	typeNotIn := "NOT IN (22, 10)"
	if CB.TypeNo == 10 {
		typeNotIn = "IN (10)"
	} else if CB.TypeNo == 22 {
		typeNotIn = "IN (22)"
	}

	/*-------------------------------------- Check If User has Booked Another Seat  --------------------------------------*/

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
		AND typeno `+typeNotIn+` AND BSEQNO != ? `, CB.Schoolno, CB.Schoolno, bookTimeStartSTR, bookTimeStartSTR, bookTimeEndSTR, bookTimeEndSTR, bookTimeStartSTR, bookTimeEndSTR, d.Bseqno).Scan(&userBookingCount)

	if userBookingCount > 0 {
		return "SE011", fmt.Errorf("user has another booking at the same time")
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
			return "SE012", fmt.Errorf("kiosk not found")
		}

		if seatKiosk.Assignable == 0 {
			return "SE013", fmt.Errorf("kiosk is not available")
		}

		d.IssueFrom = seatKiosk.Issuefrom
		d.IssueDetail = seatKiosk.Issuedetail
	}

	/*-------------------------------------- Extend Booking  --------------------------------------*/

	err := DB.Conn.Session(&gorm.Session{DryRun: false}).Transaction(func(tx *gorm.DB) error {
		var zero int64 = 0
		// var one int64 = 1
		var four int64 = 4

		if err := tx.Model(&DBmdl.SeatBooking{}).Where("bseqno = ?", d.Bseqno).Updates(map[string]interface{}{
			"useexpire":  bookTimeEndUTC,
			"printcnt":   gorm.Expr("printcnt + ?", 1),
			"extend_min": gorm.Expr("extend_min + ?", d.Minute),
			"extend_cnt": gorm.Expr("extend_cnt + ?", 1),
		}).Error; err != nil {
			return err
		}

		var extendCnt int64 = CB.ExtendCnt + 1

		if err := tx.Create(&DBmdl.SeatBookingLog{
			BSEQNO:      &d.Bseqno,
			SCHOOLNO:    &CB.Schoolno,
			SEATNO:      &CB.SeatNo,
			STATUS:      &four,
			USESTART:    &CB.Usestart,
			USEEXPIRE:   &bookTimeEndUTC,
			STARTTIME:   &CB.Starttime,
			NEWSEATNO:   &zero,
			EXTENDMIN:   &d.Minute,
			EXTENDCNT:   &extendCnt,
			LOGTIME:     &bookTimeNow,
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
