package routers

import (
	h "WISE_SOFTWARE/handlers/KIOSK"
	middlewar "WISE_SOFTWARE/middleware"

	"github.com/gin-gonic/gin"
)

func (r routes) addKiosk(rg *gin.RouterGroup) {
	kioskRouter := rg.Group("/KIOSK/")

	kioskRouter.Use(middlewar.IPFilterMiddleware())

	kioskRouter.GET("/USER/validate/:userID", h.USER_validateLogin)
	kioskRouter.GET("/sync", h.KIOSK_SYNC)
	kioskRouter.GET("/MAP/softAvailability", h.MAP_softAvailability)
	kioskRouter.GET("/MAP/deepAvailabilityBySector/:sectorNo", h.MAP_deepAvailabilitySeat)
	kioskRouter.GET("/MAP/sector_image", h.MAP_sectorImage)
	kioskRouter.GET("/SEAT/fetchSeatInfo/:seatID", h.SEAT_fetchSeatInfo)

	kioskRouter.POST("/SEAT/bookSeat", h.SEAT_bookSeat)

	kioskRouter.POST("/SEAT/extendSeat", h.SEAT_extendSeat)

	kioskRouter.POST("/SEAT/moveSeat", h.SEAT_moveSeat)

	kioskRouter.POST("/SEAT/returnSeat", h.SEAT_returnSeat)

	kioskRouter.POST("/SEAT/confirmSeat", h.SEAT_confirmSeat)

	kioskRouter.GET("/Config", h.HCConfig)
}
