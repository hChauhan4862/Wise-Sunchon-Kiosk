package function

import (
	DB "WISE_SOFTWARE/DB"
	DBmdl "WISE_SOFTWARE/DB/model"

	// "encoding/json"
	"fmt"
	"time"

	"slices"

	"gorm.io/gorm"
)

type Mdl_SeatReturn struct {
	Bseqno int64
	HcCode int64
	Issuer int64
}

func SeatReturn(d Mdl_SeatReturn) (string, error) {

	/*
		----------------------------------------MY PATTERNS FOR RETURN---------------------------------
		HC_CODE	|STATUS	|RESION	|COMMENT
		3		|	4	|	1	|	AUTO BY SYSTEM	JUST AFTER USAGE TIME END IF AUTO CONFIRMED
		2		|	4	|	2	|	AUTO BY SYSTEM	JUST AFTER USAGE TIME END IF MANUAL CONFIRMED
		1		|	6	|	9	|	AUTO BY SYSTEM	SEAT CONFIRMATION TIME OUT AROUND 30 MINUTE
		4		|	7	|	5	|	BY USER			BEFORE CONFIRMATION
		4		|	5	|	4	|	BY USER			AFTER CONFIRMATION
		5		|	7	|	10	|	BY MANAGER		IF MANAGER LET DROP THE USER
		----------------------------------------MY PATTERNS FOR RETURN---------------------------------
	*/

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

	var newStatus int64 = 0
	var newResion int64 = 0
	var isCanceled int64 = 0
	var now time.Time = time.Now()

	/* ------------------------------------------------  Get Lock  ----------------------------------------------*/
	if ok := LockGet([]string{fmt.Sprintf("BSEQNO#%v", d.Bseqno)}); !ok {
		return "LCK500", fmt.Errorf("please try in few seconds")
	}
	defer LockRelease([]string{fmt.Sprintf("BSEQNO#%v", d.Bseqno)})

	/* ---------------------------------------------- Validate Data   -------------------------------------------*/

	validHcCode := []int64{1, 2, 3, 4, 5}
	if !slices.Contains(validHcCode, d.HcCode) {
		return "SR001", fmt.Errorf("invalid status code")
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
	WHERE  bseqno = ? AND status IN ( 2, 3 ) AND iscanceled = 0`, d.Bseqno).Scan(&CB)

	if CB.BSeqNo != d.Bseqno {
		return "SR002", fmt.Errorf("booking not found or already returned")
	}

	if d.HcCode == 1 {
		if CB.Status != 2 {
			return "SR003", fmt.Errorf("booking is in use")
		}
		newStatus = 6
		newResion = 9
	} else if d.HcCode == 2 {
		if time.Now().After(CB.Useexpire) {
			TaskAlarm(CB.Useexpire.Add(2*time.Second), "BOOKING#RETURN#%v#%v", CB.BSeqNo, 2)
			return "SR004", fmt.Errorf("booking is in use")
		}
		newStatus = 4
		newResion = 2

	} else if d.HcCode == 3 {
		if time.Now().After(CB.Useexpire) {
			TaskAlarm(CB.Useexpire.Add(2*time.Second), "BOOKING#RETURN#%v#%v", CB.BSeqNo, 3)
			return "SR004", fmt.Errorf("booking is in use")
		}
		newStatus = 4
		newResion = 1

	} else if d.HcCode == 4 {
		if CB.Status == 2 {
			newStatus = 7
			newResion = 5
		}
		if CB.Status == 3 {
			newStatus = 5
			newResion = 4
		}

	} else if d.HcCode == 5 {
		newStatus = 7
		newResion = 10
	}

	if newStatus == 7 {
		isCanceled = 1
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
		// var one int64 = 1
		var log_code int64 = 5

		if err := tx.Model(&DBmdl.SeatBooking{}).Where("bseqno = ? AND STATUS IN (2,3)", d.Bseqno).Updates(map[string]interface{}{
			"STATUS":       newStatus,
			"EXPIREREASON": newResion,
			"ISCANCELED":   isCanceled,
			"RETURNTIME":   &now,
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
			STARTTIME:   &CB.Starttime,
			RETURNTIME:  &now,
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
