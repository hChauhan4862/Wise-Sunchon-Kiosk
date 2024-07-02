package KIOSK_HANDLER

import (
	f "WISE_SOFTWARE/functions"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Response_fetchSeatInfo struct {
	STATUS    int64  `json:"STATUS"`
	USE_MIN   string `json:"USE_MIN"`
	OPENTIME  string `json:"OPEN_TIME"`
	CLOSETIME string `json:"CLOSE_TIME"`
	USESTART  string `json:"USE_START"`
}

func SEAT_fetchSeatInfo(c *gin.Context) {
	// v := DB.Q.HCV_VIEW_SECTOR_USING_COUNT
	_ = c.MustGet("LIB_NO")

	seatID, _ := strconv.ParseInt(c.Params.ByName("seatID"), 10, 64)

	result, err := f.SeatInfo(seatID, time.Now())

	if err != nil || result.OPENTIME == "" {
		c.JSON(http.StatusOK, gin.H{
			"error":      true,
			"error_code": 404,
			"result":     nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":      false,
		"error_code": 200,
		"result":     result,
	})
}
