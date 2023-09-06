package logics

import (
	"net/http"
	"route-guid-api/databases"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DeleteUser(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")

	if err := db.Where("id = ?", id).Delete(&databases.User{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully!"})
}
