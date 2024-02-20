package route

import (
	"mlogreport/app/middleware"
	"mlogreport/feature/weekly/handler"
	"mlogreport/feature/weekly/repository"
	"mlogreport/feature/weekly/service"
	"mlogreport/utils/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteWeekly(c *gin.RouterGroup, db *gorm.DB) {
	weeklyRepository := repository.NewWeeklyRepository(db)
	weeklyService := service.NewWeeklyService(weeklyRepository)
	weeklyHandler := handler.NewWeeklyHandler(weeklyService)

	admin := c.Group("admin/weekly", auth.JWTMiddleware(),middleware.IsRole("advisor"))
	{
		admin.GET("/:nim",weeklyHandler.GetAllWeeklyAdvisor)
	}

	user := c.Group("weekly",auth.JWTMiddleware(),middleware.IsRole(""))
	{
		user.POST("",weeklyHandler.CreateWeekly)
		user.GET("",weeklyHandler.GetAllWeekly)
		user.PATCH("/:id",weeklyHandler.UpdateWeekly)
	}
}