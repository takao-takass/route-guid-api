package logics

import (
	"net/http"
	auth "route-guid-api/bases/auth"
	"route-guid-api/databases"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetRoutes(c *gin.Context, db *gorm.DB) {

	var userId = auth.AuthorizeUser(c, db)
	if userId < 0 {
		return
	}

	var route []databases.Route

	// Get all routes for a user.
	if err := db.Where("user_id = ?", userId).Find(&route).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, route)
}
