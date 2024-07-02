package KIOSK_HANDLER

import (
	DB "WISE_SOFTWARE/DB"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result_roomOpenCloseTimeStatus struct {
	ROOMNO       int64  `json:"ROOMNO"`
	NAME         string `json:"NAME"`
	OPENTIME     string `json:"OPENTIME"`
	CLOSETIME    string `json:"CLOSETIME"`
	Use_min      int64  `json:"USE_MIN"`
	Can_cont_min int64  `json:"CAN_CONT_MIN"`
	Cont_min     int64  `json:"CONT_MIN"`
	Max_cont_cnt int64  `json:"MAX_CONT_CNT"`
	Noshow_yn    string `json:"NOSHOW_YN"`
	Noshow_min   int64  `json:"NOSHOW_MIN"`
}
type Result_deepAvailabilityBySector struct {
	Sectorno    int64  `json:"sectorno"`
	SECTOR_NAME string `json:"SECTOR_NAME"`
	Seatno      int64  `json:"seatno"`
	NAME        string `json:"NAME"`
	Vname       string `json:"vname"`
	Status      int64  `json:"status"`
	POSX        int64  `json:"POSX"`
	POSY        int64  `json:"POSY"`
	POSW        int64  `json:"POSW"`
	POSH        int64  `json:"POSH"`
	ICONTYPE    int64  `json:"ICONTYPE"`
	Gr_min      int64  `json:"gr_min"`
	Gr_max      int64  `json:"gr_max"`
	Manx        int64  `json:"manx"`
	Many        int64  `json:"many"`
	ROOM_NAME   string `json:"ROOM_NAME"`
	USECNT      int64  `json:"USECNT"`
	XFULL       int64  `json:"XFULL"`
}

func MAP_deepAvailabilitySeat(c *gin.Context) {
	_ = c.MustGet("LIB_NO")

	var roomCode int64
	var sector_image2 string
	var typeNo int64

	DB.Conn.Raw(`SELECT roomno, TYPENO, SECTOR_IMAGE2 FROM SEAT_SECTOR WHERE SECTORNO =? AND ASSIGN_YN = 'Y'`, c.Params.ByName("sectorNo")).Row().Scan(&roomCode, &typeNo, &sector_image2)

	if roomCode == 0 {
		c.JSON(http.StatusOK, gin.H{"error": true, "error_code": "404", "result": nil})
		return
	}

	var RoomStatus Result_roomOpenCloseTimeStatus
	DB.Conn.Raw(`
	SELECT R.roomno,
       R.NAME NAME,
       R.use_min,
       R.can_cont_min,
       R.cont_min,
       R.max_cont_cnt,
       R.noshow_yn,
       R.noshow_min,
       T.opentime,
       T.closetime
FROM   seat_room2 R,
       seat_room2_ext E,
       (SELECT roomno,
               CONVERT(VARCHAR, opentime, 120)  OPENTIME,
               CONVERT(VARCHAR, closetime, 120) CLOSETIME
        FROM   dbo.Getusetimeymd(convert(varchar, getdate(), 112),?)) T
WHERE  R.roomno = E.roomno
       AND R.roomno = ?
       AND E.kiosk_yn = 'Y'
ORDER  BY sort `, roomCode, roomCode).Scan(&RoomStatus)

	var seatsResult []Result_deepAvailabilityBySector

	// 	DB.Conn.Raw(`SELECT S.seatno,
	// 	S.NAME                          NAME,
	// 	S.vname,
	// 	S.sectorno,
	// 	S.status,
	// 	S.posx2                         POSX,
	// 	S.posy2                         POSY,
	// 	S.posw,
	// 	S.posh,
	// 	S.icontype                      ICONTYPE,
	// 	S.gr_min,
	// 	S.gr_max,
	// 	S.manx,
	// 	S.many,
	// 	SE.NAME                         SECTOR_NAME,
	// 	SR.NAME                         ROOM_NAME,
	// 	(SELECT Count(*)
	// 	 FROM   seat_booking
	// 	 WHERE  seatno = S.seatno
	// 			AND ( ( usestart <= CONVERT(VARCHAR, Getdate(), 120)
	// 					AND useexpire > CONVERT(VARCHAR, Getdate(), 120) )
	// 				   OR ( usestart < CONVERT(VARCHAR, Dateadd(minute, 30,
	// 													Getdate()),
	// 								   120)
	// 						AND useexpire > CONVERT(VARCHAR, Dateadd(minute, 30,
	// 														 Getdate
	// 														 ()), 120
	// 										) )
	// 				   OR ( usestart > CONVERT(VARCHAR, Getdate(), 120)
	// 						AND useexpire <= CONVERT(VARCHAR, Dateadd(minute, 30,
	// 														  Getdate()),
	// 										 120) ) )
	// 			AND iscanceled = 0
	// 			AND status IN ( 2, 3 )) USECNT,
	// 	100                             XFULL
	// FROM   seat_seat S,
	// 	seat_sector SE,
	// 	seat_room2 SR
	// WHERE  S.sectorno = SE.sectorno
	// 	AND SE.roomno = SR.roomno
	// 	AND S.sectorno IN (SELECT sectorno
	// 					   FROM   view_sector_using_count2 C
	// 					   WHERE  C.typeno NOT IN ( 10 )
	// 							  AND C.assign_yn = ?
	// 							  AND C.libno = ?
	// 							  AND C.roomno = ?
	// 							  AND ( S.sectorno = ?
	// 									 OR (SELECT mappoint2
	// 										 FROM   seat_sector_view
	// 										 WHERE  viewsectorno = S.sectorno) IS
	// 										NULL )
	// 					  )
	// ORDER  BY seatno; `, "Y", LIB_NO, c.Params.ByName("room"), c.Params.ByName("sector")).Scan(&seatsResult)

	DB.Conn.Raw(`SELECT
		S.seatno,
		S.NAME,
		S.vname,
		S.sectorno,
		S.status,
		S.posx2 AS POSX,
		S.posy2 AS POSY,
		S.posw AS POSW,
		S.posh AS POSH,
		S.icontype,
		S.gr_min,
		S.gr_max,
		S.manx,
		S.many,
		SE.NAME AS SECTOR_NAME,
		-- SR.NAME AS ROOM_NAME,
		ISNULL(USECNT.UseCount, 0) AS USECNT,
		100 AS XFULL
	FROM seat_seat S
	JOIN seat_sector SE ON S.sectorno = SE.sectorno
	-- JOIN seat_room2 SR ON SE.roomno = SR.roomno
	LEFT JOIN (
		SELECT
			seatno,
			COUNT(*) AS UseCount
		FROM
			seat_booking
		WHERE
			iscanceled = 0
			AND status IN (2, 3)
			AND (
				(usestart <= CONVERT(VARCHAR, Getdate(), 120) AND useexpire > CONVERT(VARCHAR, Getdate(), 120))
				OR (usestart < CONVERT(VARCHAR, Dateadd(minute, 30, Getdate()), 120) AND useexpire > CONVERT(VARCHAR, Dateadd(minute, 30, Getdate()), 120))
				OR (usestart > CONVERT(VARCHAR, Getdate(), 120) AND useexpire <= CONVERT(VARCHAR, Dateadd(minute, 30, Getdate()), 120))
			)
			-- EITHER UNCOMMENT ABOVE OR BELOW TO FILTER BY TIME
		--	AND (
		--		useexpire > CONVERT(VARCHAR, Getdate(), 120)
		--	)
		GROUP BY seatno
	) USECNT ON S.seatno = USECNT.seatno
	WHERE 
		S.sectorno IN (
			SELECT sectorno FROM seat_sector
			WHERE
				ROOMNO = ?
				AND SECTOR_IMAGE2 = ?
				AND TYPENO NOT IN (10)
				AND ASSIGN_YN = ?
				AND typeNo = ?
		)

	ORDER BY seatno;`, roomCode, sector_image2, "Y", typeNo).Scan(&seatsResult)

	c.JSON(http.StatusOK, gin.H{
		"error":      false,
		"error_code": 200,
		"result": gin.H{
			"ROOM":  RoomStatus,
			"Seats": seatsResult,
		},
	})
}
