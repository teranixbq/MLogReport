package route

import (
	"mlogreport/feature/user/handler"
	"mlogreport/feature/user/repository"
	"mlogreport/feature/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteUser(c *gin.RouterGroup, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	c.POST("register",userHandler.CreateUser)

	prompt := c.Group("user")
	{
		prompt.POST("", userHandler.CreateUser)
	}
}