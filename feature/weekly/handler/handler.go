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

	result,err := weekly.weeklyService.SelectAll(nim)
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}
	c.JSON(200, helper.SuccessWithDataResponse("succes get all weekly",result))
}



