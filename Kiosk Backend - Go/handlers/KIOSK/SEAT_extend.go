package KIOSK_HANDLER

import (
	fun "WISE_SOFTWARE/functions"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type body_SEAT_extendSeat struct {
	Bseqno  int64  `json:"bseqno" binding:"required"`
	Minute  int64  `json:"minute" binding:"required"`
	TimeEnd string `json:"timeEnd" binding:"required,HHMM"`
}

func SEAT_extendSeat(c *gin.Context) {
	_ = c.MustGet("LIB_NO")
	KioskNo := c.MustGet("KIOSK_NO").(int64)

	var req body_SEAT_extendSeat
	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		return
	}

	timeEnd, _ := fun.AddTimeToTodayDate(req.TimeEnd)

	Mdl_extendSeat := fun.Mdl_SeatExtend{
		Bseqno:        req.Bseqno,
		Minute:        req.Minute,
		TimeEnd:       timeEnd,
		ISSUE_TYPE_NO: 1,
		KioskNo:       KioskNo,
	}

	status, err := fun.SeatExtend(Mdl_extendSeat)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":      true,
			"error_code": status,
			"result":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":      false,
		"error_code": nil,
		"result":     "Success",
	})

}
