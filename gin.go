package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kostya2011/dht-11-home-sensor/config"
	"github.com/kostya2011/dht-11-home-sensor/log"
)

func newGin() *gin.Engine {
	cfg := config.Values.Server

	switch {
	case cfg.Mode == "release":
		gin.SetMode(gin.ReleaseMode)
	case cfg.Mode == "test":
		gin.SetMode(gin.TestMode)
	case cfg.Mode == "debug":
		gin.SetMode(gin.DebugMode)
	default:
		log.Warn("Unknown mode selected. Falling back to `Release mode`")
		gin.SetMode(gin.ReleaseMode)
	}

	return gin.Default()
}
