package KIOSK_HANDLER

import (
	fun "WISE_SOFTWARE/functions"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HCConfig(c *gin.Context) {
	key_1 := c.Query("key_1")
	key_2 := c.Query("key_2")
	value := c.Query("value")

	valueAsInt := 0
	_, err := fmt.Sscanf(value, "%d", &valueAsInt)

	appConfig := fun.GetConfig()

	if key_1 == "timers" {
		if (err != nil) || (valueAsInt < 0) {
			c.JSON(http.StatusOK, gin.H{
				"error":      true,
				"error_code": "ER400",
				"result":     nil,
			})
			return
		}
		if key_2 == "popup" {
			appConfig.Timers.Popup = valueAsInt
		} else if key_2 == "worning" {
			appConfig.Timers.Worning = valueAsInt
		} else if key_2 == "success" {
			appConfig.Timers.Success = valueAsInt
		} else if key_2 == "error" {
			appConfig.Timers.Error = valueAsInt
		} else if key_2 == "print" {
			appConfig.Timers.Print = valueAsInt
		} else if key_2 == "floorPage" {
			appConfig.Timers.FloorPage = valueAsInt
		} else if key_2 == "bookingsDetail" {
			appConfig.Timers.BookingsDetail = valueAsInt
		} else if key_2 == "keyboard" {
			appConfig.Timers.Keyboard = valueAsInt
		} else {
			c.JSON(http.StatusOK, gin.H{
				"error":      true,
				"error_code": "ER400",
				"result":     nil,
			})
			return
		}
	} else if key_1 == "displaySectorGroupBy" {
		if value != "FLOOR" && value != "TYPE" && value != "ROOM" {
			c.JSON(http.StatusOK, gin.H{
				"error":      true,
				"error_code": "ER400",
				"result":     nil,
			})
			return
		}
		appConfig.DisplaySectorGroupBy = value
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error":      true,
			"error_code": "ER400",
			"result":     nil,
		})
		return
	}

	fun.SaveConfig(appConfig)

	c.JSON(http.StatusOK, gin.H{
		"error":      false,
		"error_code": nil,
		"result":     "OK",
	})

}
