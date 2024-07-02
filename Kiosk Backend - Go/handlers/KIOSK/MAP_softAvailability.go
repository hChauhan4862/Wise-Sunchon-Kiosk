package KIOSK_HANDLER

import (
	DB "WISE_SOFTWARE/DB"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result_softAvailability struct {
	LIBNO               int64  `json:"LIBNO"`
	Lib_name            string `json:"lib_name"`
	LIB_NAME_EN         string `json:"LIB_NAME_EN"`
	FLOORNO             int64  `json:"FLOORNO"`
	FLOOR_NAME          string `json:"FLOOR_NAME"`
	FLOOR_NAME_EN       string `json:"FLOOR_NAME_EN"`
	ROOMNO              int64  `json:"ROOMNO"`
	SECTOR_TYPE_NO      int64  `json:"SECTOR_TYPE_NO"`
	SECTOR_TYPE_NAME    string `json:"SECTOR_TYPE_NAME"`
	SECTOR_TYPE_NAME_EN string `json:"SECTOR_TYPE_NAME_EN"`
	ROOM_NAME           string `json:"ROOM_NAME"`
	ROOM_NAME_EN        string `json:"ROOM_NAME_EN"`
	SECTORNO            int64  `json:"SECTORNO"`
	SECTOR_NAME         string `json:"SECTOR_NAME"`
	SECTOR_NAME_EN      string `json:"SECTOR_NAME_EN"`
	BOOKING_YN          string `json:"BOOKING_YN"`
	ASSIGN_YN           string `json:"ASSIGN_YN"`
	TOTAL_CNT           int64  `json:"TOTAL_CNT"`
	USE_CNT             int64  `json:"USE_CNT"`
	FIX_CNT             int64  `json:"FIX_CNT"`
	FLOOR               int64  `json:"FLOOR"`
	Day_from            int64  `json:"day_from"`
	Day_to              int64  `json:"day_to"`
	Media_booking_yn    string `json:"media_booking_yn"`
	Equip_booking_yn    string `json:"equip_booking_yn"`
	FLOOR_IMAGE         string `json:"FLOOR_IMAGE"`
	SECTOR_IMAGE        string `json:"SECTOR_IMAGE"`
	TimeStart           int64  `json:"time_start"`
	TimeEnd             int64  `json:"time_end"`
	Sort                int64  `json:"sort"`
	MAPPOINT            string `json:"MAPPOINT"`
	MAPLABEL            string `json:"MAPLABEL"`
}

func MAP_softAvailability(c *gin.Context) {
	// v := DB.Q.HCV_VIEW_SECTOR_USING_COUNT
	LIB_NO := c.MustGet("LIB_NO")

	// DB.Q.HCV_VIEW_SECTOR_USING_COUNT2.SECTORNO

	var result []Result_softAvailability
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
`, "Y", LIB_NO).Scan(&result)

	// ALLOWED_ROOMS, _ := v.Where(v.TYPENO.NotIn(10), v.ASSIGNYN.Eq("Y"), v.LIBNO.Eq(LIB_NO)).Find()

	// var result []Result_softAvailability

	c.JSON(http.StatusOK, gin.H{
		"error":      false,
		"error_code": 200,
		"result":     result,
	})
}
