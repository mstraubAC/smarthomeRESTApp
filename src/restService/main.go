package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mstraubAC/smarthomeRESTApp/src/restService/accessors"
	"github.com/mstraubAC/smarthomeRESTApp/src/restService/configuration"
	"github.com/mstraubAC/smarthomeRESTApp/src/restService/controllers/aggregates"
	"github.com/mstraubAC/smarthomeRESTApp/src/restService/controllers/locations"
	"github.com/mstraubAC/smarthomeRESTApp/src/restService/middleware"
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
	logger.Debug("Configuration")
	logger.Debug(fmt.Sprintf(" - Listening on: %s", config.RestListener))
	logger.Debug(fmt.Sprintf(" - Db conn string: %.20s", config.DBUrl))

	// setup database
	databaseAccessor := accessors.DatabaseAccessor{Config: &config, Logger: logger}
	databaseAccessor.GetSqlConnection()

	// setup routing middleware
	router := gin.Default()
	router.Use(middleware.ZapLoggingHandler(logger))
	router.Use(middleware.ErrorHandler(logger))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// v1 API
	locations.RegisterRoutes(router.Group("/v1"), &config, logger, &databaseAccessor)
	aggregates.RegisterRoutes(router.Group("/v1"), &config, logger, &databaseAccessor)

	// startup router
	router.Run(config.RestListener)
}
