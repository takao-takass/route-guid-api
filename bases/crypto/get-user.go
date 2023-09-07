package bases

import (
	"route-guid-api/databases"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetUser(c *gin.Context, db *gorm.DB) int {
	token := c.Request.Header.Get("token")

	var user databases.User
	if err := db.Where("guid = ?", token).First(&user).Error; err != nil {
		return 0
	}

	return user.ID
}
