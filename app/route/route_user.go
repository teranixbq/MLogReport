package route

import (
	"mlogreport/app/middleware"
	"mlogreport/feature/user/handler"
	"mlogreport/feature/user/repository"
	"mlogreport/feature/user/service"
	"mlogreport/utils/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteUser(c *gin.RouterGroup, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	c.POST("login",userHandler.Login)

	admin := c.Group("admin/users",auth.JWTMiddleware(),middleware.IsRole("admin"))
	{
		admin.POST("", userHandler.CreateUser)
	}

	// user := c.Group("")
	// {
		
	// }
}