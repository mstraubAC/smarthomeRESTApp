package locations

import (
	"github.com/mstraubAC/smarthomeRESTApp/src/restService/accessors"
	"github.com/mstraubAC/smarthomeRESTApp/src/restService/configuration"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type handler struct {
	Config *configuration.Config
	Logger *zap.Logger
	Db     *accessors.DatabaseAccessor
}

func RegisterRoutes(router *gin.RouterGroup, config *configuration.Config, logger *zap.Logger, dbAccessor *accessors.DatabaseAccessor) {
	// setting up handler context
	h := &handler{
		Config: config,
		Logger: logger,
		Db:     dbAccessor,
	}

	// bring up sql connection
	h.Db.GetSqlConnection()

	// registering routes
	routes := router.Group("/locations")
	routes.GET("/", h.getLocations)
	routes.GET("/:locationId", h.getLocation)
}
