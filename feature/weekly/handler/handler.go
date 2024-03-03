package handler

import (
	"strings"

	"mlogreport/feature/weekly/dto/request"
	"mlogreport/feature/weekly/service"
	"mlogreport/utils/constanta"
	"mlogreport/utils/helper"

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
	data, _ := c.Get("id")
	id, _ := data.(string)

	input := request.RequestWeekly{}

	err := helper.BindJSON(c, &input)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
	}

	err = weekly.weeklyService.Insert(id, input)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR) {
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
		if strings.Contains(err.Error(), constanta.ERROR) {
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
		if strings.Contains(err.Error(), constanta.ERROR) {
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

	err := helper.BindJSON(c, &input)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
	}

	input.UsersId = nim
	err = weekly.weeklyService.UpdateWeekly(id, input)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR) {
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

	err := helper.BindJSON(c, &input)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
	}

	err = weekly.weeklyService.UpdateStatus(idUser, id, input.Status)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR) {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}
	c.JSON(200, helper.SuccessResponse("succes update status"))
}

func (weekly *weeklyhandler) CreatePeriode(c *gin.Context) {
	input := request.RequestPeriode{}

	err := helper.BindJSON(c, &input)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
		return
	}

	err = weekly.weeklyService.InsertPeriode(input)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR) {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	c.JSON(200, helper.SuccessResponse("succes create periode"))
}

func (weekly *weeklyhandler) GetAllPeriode(c *gin.Context){
	result,err := weekly.weeklyService.SelectAllPeriode()
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR) {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	if len(result) == 0 {
		c.JSON(200, helper.SuccessResponse(constanta.DATA_NULL))
		return
	}

	c.JSON(200, helper.SuccessWithDataResponse("succes get all periode",result))
}

func (weekly *weeklyhandler) UpdatePeriode(c *gin.Context) {
	id := c.Param("id")
	input := request.RequestPeriode{}

	err := helper.BindJSON(c, &input)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
		return
	}

	err = weekly.weeklyService.UpdatePeriode(id, input)
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

	c.JSON(200, helper.SuccessResponse("succes update periode"))
}
