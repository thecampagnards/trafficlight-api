package main

import (
	"net/http"
	"trafficlight/api"
	"trafficlight/usb"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(func(c *gin.Context) {
		c.Set("trafficLight", usb.New())
		c.Next()
	})

	apiGroup := r.Group("/api")
	api.TurnApiGroup(apiGroup)
	api.SwitchApiGroup(apiGroup)
	api.BlinkApiGroup(apiGroup)

	apiGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
