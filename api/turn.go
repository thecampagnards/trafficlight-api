package api

import (
	"net/http"
	"trafficlight/usb"

	"github.com/gin-gonic/gin"
)

func TurnApiGroup(api *gin.RouterGroup) {
	turn := api.Group("/turn")
	turn.DELETE("", turnOff)
	turn.GET("/red", turnOnRed)
	turn.GET("/yellow", turnOnYellow)
	turn.GET("/green", turnOnGreen)
	turn.DELETE("/red", turnOffRed)
	turn.DELETE("/yellow", turnOffYellow)
	turn.DELETE("/green", turnOffGreen)
}

func turnOnRed(c *gin.Context) {
	tl := c.MustGet("trafficLight").(usb.TrafficLight)
	err := tl.TurnOnRed()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "turnOnRed",
	})
}

func turnOnYellow(c *gin.Context) {
	tl := c.MustGet("trafficLight").(usb.TrafficLight)
	err := tl.TurnOnYellow()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "turnOnYellow",
	})
}

func turnOnGreen(c *gin.Context) {
	tl := c.MustGet("trafficLight").(usb.TrafficLight)
	err := tl.TurnOnGreen()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "turnOnGreen",
	})
}

func turnOffRed(c *gin.Context) {
	tl := c.MustGet("trafficLight").(usb.TrafficLight)
	err := tl.TurnOffRed()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "turnOffRed",
	})
}

func turnOffYellow(c *gin.Context) {
	tl := c.MustGet("trafficLight").(usb.TrafficLight)
	err := tl.TurnOffYellow()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "turnOffYellow",
	})
}

func turnOffGreen(c *gin.Context) {
	tl := c.MustGet("trafficLight").(usb.TrafficLight)
	err := tl.TurnOffGreen()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "turnOffGreen",
	})
}

func turnOff(c *gin.Context) {
	tl := c.MustGet("trafficLight").(usb.TrafficLight)
	err := tl.TurnOff()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "turnOff",
	})
}
