package main

import (
	"context"
	"fmt"

	"mlogreport/app/config"
	"mlogreport/app/database"
	"mlogreport/app/route"
	"mlogreport/app/storage"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	cfg := config.InitConfig()
	g := gin.Default()

	db := database.InitDBPostgres(cfg)
	database.DBMigration(db)
	sb := storage.InitStorage(cfg)

	g.Use(cors.Default())
	route.Run(g, db, sb)

	if cfg.MODE == "production" {
		ginLambda = ginadapter.New(g)
		lambda.Start(Handler)
	} else {
		g.Run(fmt.Sprintf(":%s", cfg.SERVERPORT))
	}
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
