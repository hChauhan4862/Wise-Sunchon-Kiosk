package KIOSK_HANDLER

import (
	fun "WISE_SOFTWARE/functions"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type body_SEAT_moveSeat struct {
	Bseqno    int64 `json:"bseqno" binding:"required"`
	NewSeatNo int64 `json:"new_seat_no" binding:"required"`
}

func SEAT_moveSeat(c *gin.Context) {
	_ = c.MustGet("LIB_NO")
	KioskNo := c.MustGet("KIOSK_NO").(int64)

	var req body_SEAT_moveSeat
	fmt.Println(c)
	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		return
	}

	Mdl_moveSeat := fun.Mdl_SeatMove{
		Bseqno:        req.Bseqno,
		New_seatNo:    req.NewSeatNo,
		ISSUE_TYPE_NO: 1,
		KioskNo:       KioskNo,
	}

	status, err := fun.SeatMove(Mdl_moveSeat)

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
