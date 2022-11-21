package locations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mstraubAC/smarhomeRESTApp/src/restService/middleware"
	"github.com/mstraubAC/smarhomeRESTApp/src/restService/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"
)

type uriParametersGetLocation struct {
	LocationId int64 `json:"locationId" uri:"locationId" binding:"required,min=0"`
}

func (h *handler) getLocation(c *gin.Context) {
	ctx := context.Background()

	// parameter validation
	params := uriParametersGetLocation{}
	if err := c.BindUri(&params); err != nil {
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
	var locations []*models.Location
	err = pgxscan.Select(ctx, sqlConn, &locations, `SELECT id, name FROM "data"."locations" WHERE id=$1 LIMIT 1`, params.LocationId)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Failed to fetch reqested data from database: %v", err))
		c.AbortWithError(http.StatusInternalServerError, &middleware.TFError{Type: middleware.ErrorSqlQueryFailed})
		return
	}

	if len(locations) == 1 {
		c.IndentedJSON(http.StatusOK, locations[0])
		return
	} else if len(locations) < 1 {
		c.AbortWithError(http.StatusNotFound, &middleware.TFError{Type: middleware.ErrorNoRecordFound})
		return
	} else if len(locations) > 1 {
		c.AbortWithError(http.StatusInternalServerError, &middleware.TFError{Type: middleware.ErrorTooManyRecordsFound})
		return
	}

}
