package KIOSK_HANDLER

import (
	DB "WISE_SOFTWARE/DB"
	fun "WISE_SOFTWARE/functions"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SeatKiosk struct {
	KioskNo           int        `gorm:"column:KioskNo"`
	Name              string     `gorm:"column:NAME"`
	EnName            string     `gorm:"column:EN_NAME"`
	IPAddr            string     `gorm:"column:IPAddr"`
	MAC               string     `gorm:"column:MAC"`
	LibNo             int64      `gorm:"column:LibNo"`
	FloorNo           int64      `gorm:"column:FloorNo"`
	IssueFrom         string     `gorm:"column:Issuefrom"`
	IssueDetail       string     `gorm:"column:IssueDetail"`
	KioskStatus       int64      `gorm:"column:KioskStatus"`
	AssignLibOnly     bool       `gorm:"column:AssignLibOnly"`
	Assignable        bool       `gorm:"column:Assignable"`
	Movable           bool       `gorm:"column:Movable"`
	Extendable        bool       `gorm:"column:Extendable"`
	Returnable        bool       `gorm:"column:Returnable"`
	StatusMemo        string     `gorm:"column:StatusMemo"`
	InsertTime        *time.Time `gorm:"column:InsertTime"`
	AdminID           string     `gorm:"column:AdminID"`
	IsDelete          bool       `gorm:"column:IsDelete"`
	PaperAmount       int64      `gorm:"column:PaperAmount"`
	PaperReplaceTime  *time.Time `gorm:"column:PaperReplaceTime"`
	PaperReplaceAdmin string     `gorm:"column:PaperReplaceAdmin"`
	PrintErrorCode    int64      `gorm:"column:PrintErrorCode"`
	PrintErrorTitle   string     `gorm:"column:PrintErrorTitle"`
	OnTime            string     `gorm:"column:OnTime"`
	OffTime           string     `gorm:"column:OffTime"`

	AllowUseRfidDevice      bool `json:"allowUseRfidDevice"`
	AllowUseVirtualKeyboard bool `json:"allowUseVirtualKeyboard"`
}

func (SeatKiosk) TableName() string {
	return "seat_kiosk2"
}

type KioskNotice struct {
	NoticeNo    int        `gorm:"column:NoticeNo"`
	NoticeType  string     `gorm:"column:NoticeType"`
	Title       string     `gorm:"column:Title"`
	ItemNo      string     `gorm:"column:ItemNo"`
	Contents    string     `gorm:"column:Contents"`
	UseStart    *time.Time `gorm:"column:UseStart"`
	UseExpire   *time.Time `gorm:"column:UseExpire"`
	InsertTime  *time.Time `gorm:"column:InsertTime"`
	AdminID     string     `gorm:"column:AdminID"`
	IsDelete    bool       `gorm:"column:IsDelete"`
	EnTitle     string     `gorm:"column:EN_TITLE"`
	EnContents  string     `gorm:"column:EN_CONTENTS"`
	NoticeColor string     `gorm:"column:NoticeColor"`
	LibNo       int        `gorm:"column:LibNo"`
	DeviceType  string     `gorm:"column:Device_Type"`
}

func (KioskNotice) TableName() string {
	return "KioskNotice"
}

func KIOSK_SYNC(c *gin.Context) {
	type Result_softAvailability struct {
		LIBNO               int64   `json:"LIBNO"`
		Lib_name            string  `json:"lib_name"`
		LIB_NAME_EN         string  `json:"LIB_NAME_EN"`
		FLOORNO             int64   `json:"FLOORNO"`
		FLOOR_NAME          string  `json:"FLOOR_NAME"`
		FLOOR_NAME_EN       string  `json:"FLOOR_NAME_EN"`
		ROOMNO              int64   `json:"ROOMNO"`
		SECTOR_TYPE_NO      int64   `json:"SECTOR_TYPE_NO"`
		SECTOR_TYPE_NAME    string  `json:"SECTOR_TYPE_NAME"`
		SECTOR_TYPE_NAME_EN string  `json:"SECTOR_TYPE_NAME_EN"`
		ROOM_NAME           string  `json:"ROOM_NAME"`
		ROOM_NAME_EN        string  `json:"ROOM_NAME_EN"`
		SECTORNO            int64   `json:"SECTORNO"`
		SECTOR_NAME         string  `json:"SECTOR_NAME"`
		SECTOR_NAME_EN      string  `json:"SECTOR_NAME_EN"`
		BOOKING_YN          string  `json:"BOOKING_YN"`
		ASSIGN_YN           string  `json:"ASSIGN_YN"`
		TOTAL_CNT           int64   `json:"TOTAL_CNT"`
		USE_CNT             int64   `json:"USE_CNT"`
		FIX_CNT             int64   `json:"FIX_CNT"`
		FLOOR               int64   `json:"FLOOR"`
		ROOM_OPENTIME       string  `json:"ROOM_OPENTIME"`
		ROOM_CLOSETIME      string  `json:"ROOM_CLOSETIME"`
		Day_from            int64   `json:"day_from"`
		Day_to              int64   `json:"day_to"`
		Media_booking_yn    string  `json:"media_booking_yn"`
		Equip_booking_yn    string  `json:"equip_booking_yn"`
		FLOOR_IMAGE         string  `json:"FLOOR_IMAGE"`
		SECTOR_IMAGE        string  `json:"SECTOR_IMAGE"`
		TimeStart           int64   `json:"time_start"`
		TimeEnd             int64   `json:"time_end"`
		Sort                int64   `json:"sort"`
		MAPPOINT            *string `json:"MAPPOINT"`
		MAPLABEL            *string `json:"MAPLABEL"`
	}

	var groupedSectors []struct {
		SECTOR_TYPE_NO      int64 `json:"SECTOR_TYPE_NO"`
		SECTOR_TYPE_NAME    string
		SECTOR_TYPE_NAME_EN string
		TOTAL_CNT           int64 `json:"TOTAL_CNT"`
		USE_CNT             int64 `json:"USE_CNT"`
		FIX_CNT             int64 `json:"FIX_CNT"`
		SECTORS             []Result_softAvailability
	}

	// v := DB.Q.HCV_VIEW_SECTOR_USING_COUNT
	LIB_NO := c.MustGet("LIB_NO")
	KioskNo := c.MustGet("KIOSK_NO").(int64)

	appConfig := fun.GetConfig()

	GROUP_SECTOR_BY := appConfig.DisplaySectorGroupBy // POSSIBLE VALUES: TYPE, ROOM, FLOOR

	// DB.Q.HCV_VIEW_SECTOR_USING_COUNT2.SECTORNO

	var sectorsData []Result_softAvailability
	DB.Conn.Raw(`SELECT	
    C.LIBNO,
    C.lib_name,
    C.LIB_EN_NAME AS LIB_NAME_EN,
    C.FLOORNO,
    C.FLOOR_NAME,
    C.FLOOR_EN_NAME AS FLOOR_NAME_EN,
    C.ROOMNO,
    C.TYPENO AS SECTOR_TYPE_NO,
	TC.NAME AS SECTOR_TYPE_NAME,
	TC.EN_NAME AS SECTOR_TYPE_NAME_EN,
    C.ROOM_NAME,
    C.ROOM_EN_NAME AS ROOM_NAME_EN,
    C.SECTORNO,
    C.SECTOR_NAME,
    C.SECTOR_EN_NAME AS SECTOR_NAME_EN,
    C.BOOKING_YN,
    C.ASSIGN_YN,
    C.TOTAL_CNT,
    C.USE_CNT,
    C.FIX_CNT,
    C.FLOOR,
	S.opentime  ROOM_OPENTIME,
	S.closetime ROOM_CLOSETIME,
    S.day_from,
    S.day_to,
    S.media_booking_yn,
    S.equip_booking_yn,
    '/MAP/KIOSK/' + FLOOR_IMAGE AS FLOOR_IMAGE,
    '/MAP/KIOSK/' + S.sector_image2 AS SECTOR_IMAGE,
    E.time_start,
    E.time_end,
    E.sort,
    SV.mapPoint2 AS MAPPOINT,
    SV.mapLabel2 AS MAPLABEL
FROM (
	select SL.NAME LIB_NAME, SL.EN_NAME LIB_EN_NAME,
		SF.NAME FLOOR_NAME, SF.EN_NAME FLOOR_EN_NAME,
		SR.NAME ROOM_NAME, SR.EN_NAME ROOM_EN_NAME,
		SE.NAME SECTOR_NAME, SE.EN_NAME SECTOR_EN_NAME,
		SE.SECTORNO, 
		SE.BOOKING_YN, SE.MOBILE_BOOKING_YN,
		SE.ASSIGN_YN, SE.MOBILE_ASSIGN_YN,
		(select count(*) from seat_seat where SECTORNO = SE.SECTORNO) TOTAL_CNT , 
		(select count(*) cnt from view_seat_booking where SECTORNO = SE.SECTORNO and STATUS in (3) AND (
			(usestart <= CONVERT(VARCHAR, Getdate(), 120) AND useexpire > CONVERT(VARCHAR, Getdate(), 120))
			OR (usestart < CONVERT(VARCHAR, Dateadd(minute, 30, Getdate()), 120) AND useexpire > CONVERT(VARCHAR, Dateadd(minute, 30, Getdate()), 120))
			OR (usestart > CONVERT(VARCHAR, Getdate(), 120) AND useexpire <= CONVERT(VARCHAR, Dateadd(minute, 30, Getdate()), 120))
		)) USE_CNT ,
		(select count(*) cnt from seat_seat where SECTORNO = SE.SECTORNO and STATUS not in (1,2)) FIX_CNT ,
		SL.LIBNO, SF.FLOORNO, SR.ROOMNO, SE.TYPENO, SF.FLOOR     
		from seat_sector SE inner join
			seat_room2 SR on SE.ROOMNO = SR.ROOMNO inner join
			seat_floor SF on SE.FLOORNO = SF.FLOORNO inner join
			seat_lib SL on SF.LIBNO = SL.LIBNO 
) C
JOIN view_seat_sector2 AS S ON S.sectorno = C.SECTORNO
JOIN seat_typecd AS TC ON TC.TYPENO = C.TYPENO
JOIN seat_room2_ext E ON S.roomno = E.roomno
LEFT JOIN seat_sector_view SV ON SV.viewsectorno = S.sectorno
JOIN seat_floor SF ON SF.FLOORNO = S.FLOORNO
WHERE C.TYPENO NOT IN (10) AND C.ASSIGN_YN = ? AND C.LIBNO = ?
ORDER BY E.sort, C.sectorno;
`, "Y", LIB_NO).Scan(&sectorsData)

	for i, s := range sectorsData {
		if s.MAPPOINT == nil {
			hasValidMapPoint := false
			// Check if there's another sector with the same SECTOR_IMAGE and valid MAPPOINT in same group
			for j, ns := range sectorsData {
				if (GROUP_SECTOR_BY == "TYPE" && s.SECTOR_TYPE_NO == ns.SECTOR_TYPE_NO ||
					GROUP_SECTOR_BY == "ROOM" && s.ROOMNO == ns.ROOMNO ||
					GROUP_SECTOR_BY == "FLOOR" && s.FLOORNO == ns.FLOORNO) && (s.SECTOR_IMAGE == ns.SECTOR_IMAGE && ns.MAPPOINT != nil) {
					hasValidMapPoint = true
					sectorsData[j].TOTAL_CNT += s.TOTAL_CNT
					sectorsData[j].USE_CNT += s.USE_CNT
					sectorsData[j].FIX_CNT += s.FIX_CNT
					break
				}
			}

			if !hasValidMapPoint {
				for _, ns := range sectorsData {
					if ns.SECTOR_IMAGE == s.SECTOR_IMAGE && ns.MAPPOINT != nil {
						sectorsData[i].MAPPOINT = ns.MAPPOINT
						sectorsData[i].MAPLABEL = ns.MAPLABEL
						break
					}
				}
			}
		}
	}

	updatedSectors := make([]Result_softAvailability, 0)
	for _, s := range sectorsData {
		if s.MAPPOINT != nil {
			updatedSectors = append(updatedSectors, s)
		}
	}

	for _, v := range updatedSectors {
		var found bool
		for i, g := range groupedSectors {
			if (GROUP_SECTOR_BY == "TYPE" && g.SECTOR_TYPE_NO == v.SECTOR_TYPE_NO) ||
				(GROUP_SECTOR_BY == "ROOM" && g.SECTOR_TYPE_NO == v.ROOMNO) ||
				(GROUP_SECTOR_BY == "FLOOR" && g.SECTOR_TYPE_NO == v.FLOORNO) {
				found = true
				groupedSectors[i].SECTORS = append(groupedSectors[i].SECTORS, v)
				groupedSectors[i].TOTAL_CNT += v.TOTAL_CNT
				groupedSectors[i].USE_CNT += v.USE_CNT
				groupedSectors[i].FIX_CNT += v.FIX_CNT
				break
			}
		}
		if !found {
			SECTOR_TYPE_NO := v.SECTOR_TYPE_NO
			SECTOR_TYPE_NAME := v.SECTOR_TYPE_NAME
			SECTOR_TYPE_NAME_EN := v.SECTOR_TYPE_NAME_EN
			if GROUP_SECTOR_BY == "ROOM" {
				SECTOR_TYPE_NO = v.ROOMNO
				SECTOR_TYPE_NAME = v.ROOM_NAME
				SECTOR_TYPE_NAME_EN = v.ROOM_NAME_EN
			} else if GROUP_SECTOR_BY == "FLOOR" {
				SECTOR_TYPE_NO = v.FLOORNO
				SECTOR_TYPE_NAME = v.FLOOR_NAME
				SECTOR_TYPE_NAME_EN = v.FLOOR_NAME_EN
			}

			groupedSectors = append(groupedSectors, struct {
				SECTOR_TYPE_NO      int64 `json:"SECTOR_TYPE_NO"`
				SECTOR_TYPE_NAME    string
				SECTOR_TYPE_NAME_EN string
				TOTAL_CNT           int64 `json:"TOTAL_CNT"`
				USE_CNT             int64 `json:"USE_CNT"`
				FIX_CNT             int64 `json:"FIX_CNT"`
				SECTORS             []Result_softAvailability
			}{
				SECTOR_TYPE_NO:      SECTOR_TYPE_NO,
				SECTOR_TYPE_NAME:    SECTOR_TYPE_NAME,
				SECTOR_TYPE_NAME_EN: SECTOR_TYPE_NAME_EN,
				TOTAL_CNT:           v.TOTAL_CNT,
				USE_CNT:             v.USE_CNT,
				FIX_CNT:             v.FIX_CNT,
				SECTORS:             []Result_softAvailability{v},
			})
		}
	}

	// ALLOWED_ROOMS, _ := v.Where(v.TYPENO.NotIn(10), v.ASSIGNYN.Eq("Y"), v.LIBNO.Eq(LIB_NO)).Find()

	var kiosk SeatKiosk
	if err := DB.Conn.Where("KioskNo = ?", KioskNo).First(&kiosk).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":      true,
			"error_code": 500,
			"result":     nil,
		})
		return
	}

	kiosk.AllowUseRfidDevice = true
	kiosk.AllowUseVirtualKeyboard = true

	var notice []KioskNotice
	if err := DB.Conn.Where("LibNo = ?", kiosk.LibNo).Find(&notice).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":      true,
			"error_code": 500,
			"result":     nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":      false,
		"error_code": 200,
		"result": gin.H{
			"config":           appConfig,
			"KioskConfig":      kiosk,
			"notice":           notice,
			"softAvailability": groupedSectors,
		},
	})
}
