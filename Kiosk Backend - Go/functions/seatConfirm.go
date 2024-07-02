package function

import (
	DB "WISE_SOFTWARE/DB"
	DBmdl "WISE_SOFTWARE/DB/model"

	// "encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Mdl_SeatConfirm struct {
	Bseqno int64
	Issuer int64
}

func SeatConfirm(d Mdl_SeatConfirm) (string, error) {
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

	var now time.Time = time.Now()
	nowSTR := now.Format("2006-01-02 15:04:05")                                   // BOOKING TIME START IN STRING FORMAT
	nowUTC, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", nowSTR+" +0000 UTC") // BOOKING TIME START IN UTC TIME

	/* ------------------------------------------------  Get Lock  ----------------------------------------------*/
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
	use_min,
	(SELECT closetime
	 FROM   dbo.Getusetimeymd(CONVERT(VARCHAR, V.useexpire, 112), V.roomno))
									 CLOSETIME
	FROM   view_seat_booking V
	WHERE  bseqno = ? AND status IN ( 2 ) AND iscanceled = 0`, d.Bseqno).Scan(&CB)

	if CB.BSeqNo != d.Bseqno {
		return "SC001", fmt.Errorf("booking not found or already confirmed")
	}

	if CB.Useexpire.Before(nowUTC) {
		return "SC002", fmt.Errorf("booking expired")
	}

	/* ------------------------------------ GET ISSUER INFORMATION  ----------------------------------- */

	var IssueFrom int64
	var IssueDetail string

	if d.Issuer == 0 {
		IssueFrom = 3001
		IssueDetail = "0:0:0:0:0:0:0:1"
	} else if d.Issuer == 2 {
		IssueFrom = 1
		IssueDetail = "pc"
	} else if d.Issuer == 3 {
		IssueFrom = 10
		IssueDetail = "mobile"
	} else {
		var seatKiosk seatKiosk
		DB.Conn.Raw(`SELECT kioskno,
			issuefrom,
			issuedetail,
			LibNo,
			AssignLibOnly,
			Assignable
		FROM   seat_kiosk2
		WHERE  kioskno = ?`, d.Issuer).Scan(&seatKiosk)

		if seatKiosk.Kioskno == 0 {
			return "SR005", fmt.Errorf("could not find issuer")
		}

		IssueFrom = seatKiosk.Issuefrom
		IssueDetail = seatKiosk.Issuedetail
		d.Issuer = 1
	}

	/*-------------------------------------- Return Booking  --------------------------------------*/

	err := DB.Conn.Session(&gorm.Session{DryRun: false}).Transaction(func(tx *gorm.DB) error {
		var zero int64 = 0
		var log_code int64 = 2
		var newStatus int64 = 3

		if err := tx.Model(&DBmdl.SeatBooking{}).Where("bseqno = ? AND STATUS IN (2,3)", d.Bseqno).Updates(map[string]interface{}{
			"STATUS":    newStatus,
			"STARTTIME": now,
		}).Error; err != nil {
			return err
		}

		if err := tx.Create(&DBmdl.SeatBookingLog{
			BSEQNO:      &d.Bseqno,
			SCHOOLNO:    &CB.Schoolno,
			SEATNO:      &CB.SeatNo,
			STATUS:      &log_code,
			USESTART:    &CB.Usestart,
			USEEXPIRE:   &CB.Useexpire,
			STARTTIME:   &now,
			NEWSEATNO:   &zero,
			EXTENDMIN:   &zero,
			EXTENDCNT:   &zero,
			LOGTIME:     &now,
			ISSUEFROM:   &IssueFrom,
			ISSUEDETAIL: &IssueDetail,
			ISSUETYPENO: &d.Issuer,
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
