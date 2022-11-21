package aggregates

import (
	"context"
	"mstraubAC/smarthome-restService/configuration"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type handler struct {
	DB     *pgxpool.Pool
	Config *configuration.Config
	Logger *zap.Logger
}

func RegisterRoutes(router *gin.RouterGroup, config *configuration.Config, logger *zap.Logger) {
	// setting up handler context
	h := &handler{
		Config: config,
		Logger: logger,
	}

	// bring up sql connection
	h.GetSqlConnection()

	// registering routes
	routes := router.Group("/aggregates")
	routes.GET("/heatpumpdaily", h.getHeatpumpAggregatesDaily)
	// routes.GET("/:locationId", h.getLocation)
}

func (h *handler) GetSqlConnection() (*pgxpool.Pool, error) {
	if h.DB == nil {
		h.Logger.Info("No PostgreSQL connection available. Creating new connection pool")
		dbConn, err := pgxpool.New(context.Background(), h.Config.DBUrl)
		if err != nil {
			h.Logger.Error("PostgresSQL connection could not be established: " + err.Error())
			return h.DB, err
		}
		h.DB = dbConn
	}

	return h.DB, nil
}
