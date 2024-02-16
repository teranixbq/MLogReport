package main

import (
	"fmt"
	"mlogreport/app/config"
	"mlogreport/app/database"
	"mlogreport/app/route"
	"mlogreport/app/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	cfg := config.InitConfig()
	db := database.InitDBPostgres(cfg)
	database.DBMigration(db)
	sb := storage.InitStorage(cfg)

	g.Use(cors.Default())
	route.Run(g, db, sb)

	g.Run(fmt.Sprintf(":%s", cfg.SERVERPORT))
}
