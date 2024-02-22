package handler

import (
	"mlogreport/feature/weekly/dto/request"
	"mlogreport/feature/weekly/service"
	"mlogreport/utils/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

type weeklyhandler struct {
	weeklyService service.WeeklyServiceInterface
}

func NewWeeklyHandler(weeklyService service.WeeklyServiceInterface) *weeklyhandler {
	return &weeklyhandler{
		weeklyService: weeklyService,
	}
}

func (weekly *weeklyhandler) CreateWeekly(c *gin.Context) {
	id, _ := c.Get("id")
	nim, _ := id.(string)

	input := request.RequestWeekly{}

	err := c.Bind(&input)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
	}

	err = weekly.weeklyService.Insert(nim, input)
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}
	c.JSON(200, helper.SuccessResponse("succes insert weekly log"))
}

func (weekly *weeklyhandler) GetAllWeekly(c *gin.Context) {
	id, _ := c.Get("id")
	nim, _ := id.(string)

	result, err := weekly.weeklyService.SelectAll(nim)
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}
	c.JSON(200, helper.SuccessWithDataResponse("succes get all weekly", result))
}

func (weekly *weeklyhandler) GetAllWeeklyAdvisor(c *gin.Context) {
	id, _ := c.Get("id")
	nip, _ := id.(string)
	nim := c.Param("nim")

	result, err := weekly.weeklyService.SelectAllWeeklyAdvisor(nip, nim)
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}
	c.JSON(200, helper.SuccessWithDataResponse("succes get all weekly", result))
}

func (weekly *weeklyhandler) UpdateWeekly(c *gin.Context) {
	data, _ := c.Get("id")
	nim, _ := data.(string)
	id := c.Param("id")
	input := request.RequestWeekly{}

	err := c.Bind(&input)
	if err != nil {
		if err != nil {
			c.JSON(400, helper.ErrorResponse(err.Error()))
		}
	}

	input.UsersId = nim
	err = weekly.weeklyService.UpdateWeekly(id, input)
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}
	c.JSON(200, helper.SuccessResponse("succes update weekly log"))
}

func (weekly *weeklyhandler) UpdateStatus(c *gin.Context) {
	idUser := c.Param("iduser")
	id := c.Param("id")
	input := request.RequestStatus{}

	err := c.Bind(&input)
	if err != nil {
		if err != nil {
			c.JSON(400, helper.ErrorResponse(err.Error()))
		}
	}

	err = weekly.weeklyService.UpdateStatus(idUser, id, input.Status)
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}
	c.JSON(200, helper.SuccessResponse("succes update status"))
}
