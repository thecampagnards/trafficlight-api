package api

import (
	"net/http"
	"trafficlight/usb"

	"github.com/gin-gonic/gin"
)

func BlinkApiGroup(api *gin.RouterGroup) {
	blink := api.Group("/blink")
	blink.DELETE("", turnOff)
	blink.GET("/red", BlinkRed)
	blink.GET("/yellow", BlinkYellow)
	blink.GET("/green", BlinkGreen)
}

func BlinkRed(c *gin.Context) {
	tl := c.MustGet("trafficLight").(usb.TrafficLight)
	err := tl.BlinkRed()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "BlinkRed",
	})
}

func BlinkYellow(c *gin.Context) {
	tl := c.MustGet("trafficLight").(usb.TrafficLight)
	err := tl.BlinkYellow()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "BlinkYellow",
	})
}

func BlinkGreen(c *gin.Context) {
	tl := c.MustGet("trafficLight").(usb.TrafficLight)
	err := tl.BlinkGreen()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "BlinkGreen",
	})
}
