package aggregates

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
	routes := router.Group("/aggregates")
	routes.GET("/heatpump/daily", h.getHeatpumpAggregatesDaily)
	routes.GET("/heatpump/monthly", h.getHeatpumpAggregatesMonthly)
	routes.GET("/heatpump/yearly", h.getHeatpumpAggregatesYearly)
	routes.GET("/electricconsumption/moneyflow/daily", h.getElectricEnergyMoneyFlowDaily)
	routes.GET("/electricconsumption/moneyflow/monthly", h.getElectricEnergyMoneyFlowMonthly)
	routes.GET("/electricconsumption/moneyflow/yearly", h.getElectricEnergyMoneyFlowYearly)
	routes.GET("/electricconsumption/flow/daily", h.getElectricEnergyFlowDaily)
	routes.GET("/electricconsumption/flow/monthly", h.getElectricEnergyFlowMonthly)
	routes.GET("/electricconsumption/flow/yearly", h.getElectricEnergyFlowYearly)
	// routes.GET("/:locationId", h.getLocation)
}
