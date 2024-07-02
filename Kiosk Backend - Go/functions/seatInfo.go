package function

import (
	DB "WISE_SOFTWARE/DB"
	"time"
)

type mdl_SeatInfo struct {
	STATUS    int64  `json:"STATUS"`
	USE_MIN   string `json:"USE_MIN"`
	OPENTIME  string `json:"OPEN_TIME"`
	CLOSETIME string `json:"CLOSE_TIME"`
	USESTART  string `json:"USE_START"`
}

func SeatInfo(seatID int64, t time.Time) (mdl_SeatInfo, error) {
	var result mdl_SeatInfo
	err := DB.Conn.Raw(`SELECT	S.status,
	S.use_min,
	Se.opentime,
	Se.closetime,
		(SELECT TOP 1 CONVERT(VARCHAR, usestart, 120) USESTART
			FROM   view_seat_booking
				WHERE  seatno = S.seatno
					AND usestart >= ?
					AND status IN ( 2, 3 )
				ORDER  BY usestart
		) USESTART
	FROM   view_seat_seat S,
			view_seat_sector Se
	WHERE  S.seatno = ?
			AND Se.sectorno = S.sectorno `, FormatTime(t, 120), seatID).Scan(&result).Error

	return result, err
}
