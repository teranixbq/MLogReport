package handler

import (
	"strings"

	"mlogreport/feature/user/dto/request"
	"mlogreport/feature/user/service"
	"mlogreport/utils/constanta"
	"mlogreport/utils/helper"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserServiceInterface
}

func NewUserHandler(userService service.UserServiceInterface) *userHandler {
	return &userHandler{userService: userService}
}

func (user *userHandler) CreateUser(c *gin.Context) {
	input := request.RequestUser{}

	err := helper.BindJSON(c, &input)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
		return
	}

	err = user.userService.InsertUser(input)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR) {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	c.JSON(201, helper.SuccessResponse("success insert data"))
}

func (user *userHandler) Login(c *gin.Context) {
	input := request.RequestLogin{}
	err := helper.BindJSON(c, &input)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
		return
	}

	result, err := user.userService.Login(input)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR) {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		if strings.Contains(err.Error(), constanta.NOT_FOUND) {
			c.AbortWithStatusJSON(404, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	c.JSON(200, helper.SuccessWithDataResponse("success login", result))
}

func (user *userHandler) GetProfile(c *gin.Context) {
	id, _ := c.Get("id")
	nim, _ := id.(string)

	result, err := user.userService.SelectUserById(nim)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR) {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		if strings.Contains(err.Error(), constanta.NOT_FOUND) {
			c.AbortWithStatusJSON(404, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	c.JSON(200, helper.SuccessWithDataResponse("succes get profile", result))
}

func (user *userHandler) UpdateProfile(c *gin.Context) {
	id, _ := c.Get("id")
	nim, _ := id.(string)
	input := request.RequestUpdateProfile{}

	err := helper.BindJSON(c, &input)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
		return
	}

	err = user.userService.UpdateProfile(nim, input)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR) {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		if strings.Contains(err.Error(), constanta.NOT_FOUND) {
			c.AbortWithStatusJSON(404, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	c.JSON(200, helper.SuccessResponse("succes update profile"))
}
