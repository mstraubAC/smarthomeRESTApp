package locations

import (
	"context"
	"mstraubAC/smarthome-restService/models"
	"net/http"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"
)

// for testing keep locations in memory
// var locations = []models.Location{
// 	{Id: 1, Name: "Test location"},
// }

// GetLocations godoc
// @Summary		lists all locations
// @Description	list all locations
// @Produce		json
// @Success		200 	{object}	[]models.Location
// @Router		/locations [get]
func (h *handler) getLocations(c *gin.Context) {
	ctx := context.Background()

	sqlConn, err := h.GetSqlConnection()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var locations []*models.Location
	pgxscan.Select(ctx, sqlConn, &locations, `SELECT id, name FROM "data"."locations"`)
	if len(locations) > 0 {
		c.IndentedJSON(http.StatusOK, locations)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
