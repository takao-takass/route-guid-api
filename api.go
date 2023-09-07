package main

import (
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"route-guid-api/configurations"
	"route-guid-api/routes"
)

var db *sql.DB

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var config configurations.Config
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: config.DatabaseServerName,
	})

	db, err := gorm.Open("mysql", config.DatabaseConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()

	routes.UserRoute(r, db)
	routes.RouteRoute(r, db)

	r.Run(":8000")
}
