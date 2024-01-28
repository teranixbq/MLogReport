package handler

import (
	"mlogreport/feature/user/dto/request"
	"mlogreport/feature/user/service"
	"mlogreport/utils/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserServiceInterface
}

func NewUserHandler(userService service.UserServiceInterface) *userHandler{
	return &userHandler{userService: userService}
}

func (user *userHandler) CreateUser(c *gin.Context) {
	input := request.RequestUser{}
	err := c.Bind(&input)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
		return
	}

	err = user.userService.InsertUser(input)
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}
	
	c.JSON(201, helper.SuccessResponse("success insert data"))
}