package KIOSK_HANDLER

import (
	fun "WISE_SOFTWARE/functions"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type body_SEAT_returnSeat struct {
	Bseqno int64 `json:"bseqno" binding:"required"`
}

func SEAT_returnSeat(c *gin.Context) {
	_ = c.MustGet("LIB_NO")
	KioskNo := c.MustGet("KIOSK_NO").(int64)

	var req body_SEAT_returnSeat
	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		return
	}

	Mdl_returnSeat := fun.Mdl_SeatReturn{
		Bseqno: req.Bseqno,
		Issuer: KioskNo,
		HcCode: 4,
	}

	status, err := fun.SeatReturn(Mdl_returnSeat)

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
