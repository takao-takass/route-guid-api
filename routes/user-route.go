package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	userLogics "route-guid-api/logics/user"
)

func UserRoute(r *gin.Engine, db *sql.DB) {
	// hello world
	r.GET("/", userLogics.HelloWorld)

	// Create
	r.POST("/users", func(c *gin.Context) {
		userLogics.CreateUser(c, db)
	})

	// Read
	r.GET("/users/:id", func(c *gin.Context) {
		userLogics.ReadUser(c, db)
	})

	// Update
	r.PUT("/users/:id", func(c *gin.Context) {
		userLogics.UpdateUser(c, db)
	})

	// Delete
	r.DELETE("/users/:id", func(c *gin.Context) {
		userLogics.DeleteUser(c, db)
	})
}
