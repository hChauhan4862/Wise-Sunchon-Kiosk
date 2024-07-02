package KIOSK_HANDLER

import (
	fun "WISE_SOFTWARE/functions"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MAP_sectorImage(c *gin.Context) {
	m := c.Query("map")
	image := fun.SECTOR_IMAGE_TO_BASE64(m)

	c.JSON(http.StatusOK, gin.H{
		"error":      false,
		"error_code": nil,
		"result":     image,
	})

}
