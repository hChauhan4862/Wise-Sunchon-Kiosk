package KIOSK_HANDLER

import (
	fun "WISE_SOFTWARE/functions"
	"net/http"

	// "time"

	"github.com/gin-gonic/gin"
)

type body_SEAT_bookSeat struct {
	UserID    string `json:"userID" binding:"required"`
	SeatNo    int64  `json:"seatNo" binding:"required"`
	TimeStart string `json:"timeStart" binding:"required,HHMM"`
	TimeEnd   string `json:"timeEnd" binding:"required,HHMM"`
}

func SEAT_bookSeat(c *gin.Context) {
	_ = c.MustGet("LIB_NO")

	KIOSK_NO := c.MustGet("KIOSK_NO").(int64)

	var req body_SEAT_bookSeat
	if err := c.BindJSON(&req); err != nil {
		return
	}

	timeStart, _ := fun.AddTimeToTodayDate(req.TimeStart) // Booking Should be start from as per user request

	// timeStart := time.Now() // Booking Should be start from now
	// timeStart, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2023-08-11 14:50:00 +0900 KST") // Bypass Booking Start Time

	timeEnd, _ := fun.AddTimeToTodayDate(req.TimeEnd)

	Mdl_BookSeat := fun.Mdl_SeatBook{
		SeatNo:        req.SeatNo,
		Schoolno:      req.UserID,
		TimeStart:     timeStart,
		TimeEnd:       timeEnd,
		ISSUE_TYPE_NO: 1,
		KioskNo:       KIOSK_NO,
	}

	status, err := fun.SeatBook(Mdl_BookSeat)

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
		"result": gin.H{
			"bseqno": status,
		},
	})

}
