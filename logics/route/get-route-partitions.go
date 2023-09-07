package logics

import (
	"net/http"
	auth "route-guid-api/bases/auth"
	"route-guid-api/databases"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetRoutePatitions(c *gin.Context, db *gorm.DB) {

	var userId = auth.AuthorizeUser(c, db)
	if userId < 0 {
		return
	}

	routeId := c.Param("routeId")

	var route databases.Route

	// Find route by user id and route id.
	if err := db.Where("user_id = ?", userId).First(&route).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var routePartition []databases.RoutePartition

	// Get all route partitions for a route.
	if err := db.Where("route_id = ?", routeId).Find(&routePartition).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, routePartition)
}
