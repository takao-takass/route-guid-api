package bases

import (
	"net/http"
	crypt "route-guid-api/bases/crypto"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func AuthorizeUser(c *gin.Context, db *gorm.DB) int {

	if userId := crypt.GetUser(c, db); userId > 0 {
		return userId
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	return -1
}
