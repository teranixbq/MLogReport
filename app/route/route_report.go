package route

import (
	"mlogreport/app/middleware"
	"mlogreport/app/storage"
	"mlogreport/feature/report/handler"
	"mlogreport/feature/report/repository"
	"mlogreport/feature/report/service"
	"mlogreport/utils/auth"

	supabase "github.com/supabase-community/storage-go"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteReport(c *gin.RouterGroup, db *gorm.DB, sb *supabase.Client) {
	supabaseConfig := storage.NewStorage(sb)
	reportRepository := repository.NewReportRepository(db,supabaseConfig)
	reportService := service.NewReportService(reportRepository)
	reportHandler := handler.NewReportHandler(reportService)

	admin := c.Group("admin/report",auth.JWTMiddleware(),middleware.IsRole("advisor"))
	{
		admin.GET("",reportHandler.GetAllReport)
	}

	user := c.Group("report",auth.JWTMiddleware(),middleware.IsRole(""))
	{
		user.POST("",reportHandler.InsertUpdate)
		user.GET("",reportHandler.GetReport)
	}
}