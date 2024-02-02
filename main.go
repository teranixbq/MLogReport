package main

import (
	"fmt"
	"mlogreport/app/config"
	"mlogreport/app/database"
	"mlogreport/app/route"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	cfg := config.InitConfig()
	db := database.InitDBPostgres(cfg)
	database.DBMigration(db)

	g.Use(cors.Default())
	route.Run(g, db)

	g.Run(fmt.Sprintf(":%s", cfg.SERVERPORT))
}
