package handler

import (
	"mlogreport/feature/admin/dto/request"
	"mlogreport/feature/admin/service"
	"mlogreport/utils/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

type adminHandler struct {
	adminService service.AdminServiceInterface
}

func NewAdminHandler(adminService service.AdminServiceInterface) *adminHandler {
	return &adminHandler{
		adminService: adminService,
	}
}

func (admin *adminHandler) CreateAdvisor(c *gin.Context) {
	input := request.CreateAdvisor{}

	// _, role, errAuth := auth.ExtractToken(c)
	// if errAuth != nil {
	// 	c.JSON(400, helper.ErrorResponse(errAuth.Error()))
	// 	return
	// }

	// if role != "admin" {
	// 	c.JSON(400, helper.ErrorResponse("error : you don't have access"))
	// 	return
	// }

	err := c.Bind(&input)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
		return
	}

	err = admin.adminService.CreateAdvisor(input)
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

func (admin *adminHandler) Login(c *gin.Context) {
	input := request.AdminLogin{}

	err := c.Bind(&input)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
		return
	}

	response, err := admin.adminService.Login(input)
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	c.JSON(200, helper.SuccessWithDataResponse("success login", response))
}
