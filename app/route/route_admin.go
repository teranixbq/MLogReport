package route

import (
	"mlogreport/feature/admin/handler"
	"mlogreport/feature/admin/repository"
	"mlogreport/feature/admin/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteAdmin(c *gin.RouterGroup, db *gorm.DB) {
	adminRepository := repository.NewPromptRepository(db)
	adminService := service.NewAdminService(adminRepository)
	adminHandler := handler.NewAdminHandler(adminService)

	prompt := c.Group("admin")
	{
		prompt.POST("", adminHandler.CreateAdvisor)
	}
}