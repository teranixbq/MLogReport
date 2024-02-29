package handler

import (
	"strings"

	"mlogreport/feature/admin/dto/request"
	"mlogreport/feature/admin/service"
	"mlogreport/utils/constanta"
	"mlogreport/utils/helper"

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

	err := c.Bind(&input)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
		return
	}

	err = admin.adminService.CreateAdvisor(input)
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

func (admin *adminHandler) Login(c *gin.Context) {
	input := request.AdminLogin{}

	err := c.Bind(&input)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
		return
	}

	response, err := admin.adminService.Login(input)
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

	c.JSON(200, helper.SuccessWithDataResponse("success login", response))
}

func (admin *adminHandler) GetAllAdvisor(c *gin.Context) {
	data := c.MustGet("pagination").(gin.H)
	page := data["page"].(int)
	limit := data["limit"].(int)

	result, metaInfo, err := admin.adminService.SelectAllAdvisor(page, limit)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR) {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	c.JSON(200, helper.SuccessWithPageResponse("success get all advisor", metaInfo, result))
}

func (admin *adminHandler) GetAdvisor(c *gin.Context) {
	id := c.Param("id")

	result, err := admin.adminService.SelectAdvisor(id)
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

	c.JSON(200, helper.SuccessWithDataResponse("succes get advisor", result))
}

func (admin *adminHandler) CreateListColleges(c *gin.Context) {
	input := request.ListCollege{}

	err := c.Bind(&input)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
		return
	}

	err = admin.adminService.InsertList(input)
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

	c.JSON(201, helper.SuccessResponse("success add college"))
}

func (admin *adminHandler) DeleteAdvisor(c *gin.Context) {
	id := c.Param("id")

	err := admin.adminService.DeleteAdvisor(id)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR) {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	c.JSON(200, helper.SuccessResponse("success delete advisor"))
}
