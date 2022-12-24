package aggregates

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mstraubAC/smarthomeRESTApp/src/restService/middleware"
	"github.com/mstraubAC/smarthomeRESTApp/src/restService/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"
)

func (h *handler) getElectricEnergyFlowMonthly(c *gin.Context) {
	ctx := context.Background()

	// get sql accessor
	sqlConn, err := h.Db.GetSqlConnection()
	if err != nil {
		// TODO: Unify logging format to include tracing Ids
		h.Logger.Error(fmt.Sprintf("Failed to fetch a database connection: %v", err))
		c.AbortWithError(http.StatusInternalServerError, &middleware.TFError{Type: middleware.ErrorNoDatabaseConnection})
		return
	}

	// perform action
	var locations []*models.ElectricEnergyFlow
	err = pgxscan.Select(ctx, sqlConn, &locations,
		`SELECT 
			logdate
			,electricConsumeValue
			,electricGridFeedInValue
			,electricGridPvGenerationValue
			,electricHeatingConsumeValue
			,electricItConsumeValue
			,electricWallboxConsumeValue
		FROM aggregation."vElectricEnergyFlowsMonthly" 
		ORDER BY logdate ASC`)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Failed to fetch reqested data from database: %v", err))
		println(fmt.Sprintf("Failed to fetch reqested data from database: %v", err))
		c.AbortWithError(http.StatusInternalServerError, &middleware.TFError{Type: middleware.ErrorSqlQueryFailed})
		return
	}

	if len(locations) > 0 {
		c.IndentedJSON(http.StatusOK, locations)
		return
	} else if len(locations) < 1 {
		c.AbortWithError(http.StatusNotFound, &middleware.TFError{Type: middleware.ErrorNoRecordFound})
		return
	}

}
