package main

import (
	"mstraubAC/smarthome-restService/accessors"
	"mstraubAC/smarthome-restService/configuration"
	"mstraubAC/smarthome-restService/controllers/aggregates"
	"mstraubAC/smarthome-restService/controllers/locations"
	"mstraubAC/smarthome-restService/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// logger
	logger, _ := zap.NewDevelopment()

	// configuration
	config, err := configuration.LoadConfig(*logger)
	if err != nil {
		logger.Fatal("Failed to read config: " + err.Error())
		logger.Fatal("Terminating")
		return
	}

	// setup database
	databaseAccessor := accessors.DatabaseAccessor{Config: &config, Logger: logger}

	// setup routing middleware
	router := gin.Default()
	router.Use(middleware.ZapLoggingHandler(logger))
	router.Use(middleware.ErrorHandler(logger))

	// register routes
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(200, config)
	})

	// v1 API
	locations.RegisterRoutes(router.Group("/v1"), &config, logger)
	aggregates.RegisterRoutes(router.Group("/v1"), &config, logger, &databaseAccessor)

	// startup router
	router.Run(config.RestListener)
}
