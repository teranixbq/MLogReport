package route

import (
	"mlogreport/app/middleware"
	"mlogreport/feature/admin/handler"
	"mlogreport/feature/admin/repository"
	"mlogreport/feature/admin/service"
	"mlogreport/utils/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteAdmin(c *gin.RouterGroup, db *gorm.DB) {
	adminRepository := repository.NewPromptRepository(db)
	adminService := service.NewAdminService(adminRepository)
	adminHandler := handler.NewAdminHandler(adminService)

	c.POST("/admin/login", adminHandler.Login)

	admin := c.Group("admin", auth.JWTMiddleware(), middleware.IsRole("admin"))
	{
		admin.POST("", adminHandler.CreateAdvisor)
		admin.POST("/add",adminHandler.CreateListColleges)
		admin.DELETE("/:id",adminHandler.DeleteAdvisor)
	}
}
