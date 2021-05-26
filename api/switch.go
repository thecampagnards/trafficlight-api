package api

import (
	"net/http"
	"trafficlight/usb"

	"github.com/gin-gonic/gin"
)

func SwitchApiGroup(api *gin.RouterGroup) {
	switchGroup := api.Group("/switch")
	switchGroup.DELETE("", turnOff)
	switchGroup.GET("/red", switchRed)
	switchGroup.GET("/yellow", switchYellow)
	switchGroup.GET("/green", switchGreen)
}

func switchRed(c *gin.Context) {
	tl := c.MustGet("trafficLight").(usb.TrafficLight)
	err := tl.SwitchRed()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "switchRed",
	})
}

func switchYellow(c *gin.Context) {
	tl := c.MustGet("trafficLight").(usb.TrafficLight)
	err := tl.SwitchYellow()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "switchYellow",
	})
}

func switchGreen(c *gin.Context) {
	tl := c.MustGet("trafficLight").(usb.TrafficLight)
	err := tl.SwitchGreen()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "switchGreen",
	})
}
