package middleware

import (
	DB "WISE_SOFTWARE/DB"

	"github.com/gin-gonic/gin"
)

type seatKiosk struct {
	Kioskno       int64  `json:"kioskno"`
	Issuefrom     int64  `json:"issuefrom"`
	Issuedetail   string `json:"issuedetail"`
	LibNo         int64  `json:"LibNo"`
	AssignLibOnly int64  `json:"AssignLibOnly"`
	Assignable    int64  `json:"Assignable"`
}

func IPFilterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the client's IP address

		clientIP := c.ClientIP()

		// Check if the IP address starts with "192."
		// if !strings.HasPrefix(clientIP, "192.") && clientIP != "::1" {
		// 	c.String(http.StatusForbidden, "404 page not found")
		// 	c.Abort()
		// 	return
		// }

		// clientIP = "103.83.145.146"

		var seatKiosk seatKiosk

		DB.Conn.Raw(`SELECT kioskno,
			issuefrom,
			issuedetail,
			LibNo,
			AssignLibOnly,
			Assignable
		FROM   seat_kiosk2
		WHERE  IPAddr = ?`, clientIP).Scan(&seatKiosk)

		if seatKiosk.Issuefrom == 0 {
			seatKiosk.Kioskno = 1
			seatKiosk.LibNo = 1

			// c.String(http.StatusForbidden, "404 page not found")
			// c.Abort()
			// return
		}

		var kioskno int64 = seatKiosk.Kioskno
		var LIB_NO int64 = seatKiosk.LibNo
		c.Set("LIB_NO", LIB_NO)
		c.Set("KIOSK_NO", kioskno)

		c.Next()
	}
}
