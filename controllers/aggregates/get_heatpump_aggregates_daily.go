package aggregates

import (
	"context"
	"fmt"
	"mstraubAC/smarthome-restService/middleware"
	"mstraubAC/smarthome-restService/models"
	"net/http"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"
)

type queryParametersGetHeatpumpAggregationDaily struct {
	StartDate time.Time `json:"startDate" form:"startDate" time_format:"2006-01-02" binding:"required"`
	EndDate   time.Time `json:"endDate" form:"endDate" time_format:"2006-01-02" binding:"required,gtfield=StartDate"`
}

func (h *handler) getHeatpumpAggregatesDaily(c *gin.Context) {
	ctx := context.Background()

	// parameter validation
	params := queryParametersGetHeatpumpAggregationDaily{}
	if err := c.BindQuery(&params); err != nil {
		// TODO: better reporting of validation errors but machine parsable, like https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/
		c.AbortWithError(http.StatusBadRequest, &middleware.TFError{Type: middleware.ErrorRequestParameterInvalid, Detail: fmt.Sprintf("%v", err)})
		return
	}

	// get sql accessor
	sqlConn, err := h.Db.GetSqlConnection()
	if err != nil {
		// TODO: Unify logging format to include tracing Ids
		h.Logger.Error(fmt.Sprintf("Failed to fetch a database connection: %v", err))
		c.AbortWithError(http.StatusInternalServerError, &middleware.TFError{Type: middleware.ErrorNoDatabaseConnection})
		return
	}

	// perform action
	var locations []*models.HeatpumpPowerAggregatesType
	err = pgxscan.Select(ctx, sqlConn, &locations,
		`SELECT 
				logdate, tagesarbeitszahlinclcontrolandpumps, tagesarbeitszahlsolvis, tagesarbeitszahlfullelectricmeasurement, 
				totalelectricenergy, totalelectricenergyfullmeasurement, totalelectricenergysolvismeasurement, totalheatingenergy,
				heatpumpthermalpowerenergy, heatpumpresistanceheatingenergy, outsidetemperatureavg, flowtemperaturecircuit1avg 
			FROM "data"."fnGetHeatingKpisDaily"($1,$2)`, params.StartDate, params.EndDate)
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
