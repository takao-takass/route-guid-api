package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	routeLogics "route-guid-api/logics/route"
)

func RouteRoute(r *gin.Engine, db *gorm.DB) {

	// Get all routes.
	r.GET("/routes", func(c *gin.Context) {
		routeLogics.GetRoutes(c, db)
	})

	// Get all route partitions for a route.
	r.GET("/routes/:routeId/partitions", func(c *gin.Context) {
		routeLogics.GetRoutePatitions(c, db)
	})

	// Get all route paths for a route partition.
	r.GET("/routes/partitions/:partitionId/paths", func(c *gin.Context) {
		routeLogics.GetRoutePaths(c, db)
	})

}
