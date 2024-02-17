package handler

import (
	"mlogreport/feature/report/dto/request"
	"mlogreport/feature/report/service"
	"mlogreport/utils/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

type reportHandler struct {
	reportService service.ReportServiceInterface
}

func NewReportHandler(reportService service.ReportServiceInterface) *reportHandler {
	return &reportHandler{
		reportService: reportService,
	}
}

func (report *reportHandler) InsertUpdate(c *gin.Context) {
	data, _ := c.Get("id")
	nim, _ := data.(string)

	finalReport, _ := c.FormFile("final_report")
	transcript, _ := c.FormFile("transcript")
	certification, _ := c.FormFile("certification")

	input := request.MultipartToRequestReport(finalReport, transcript, certification)
	 err := report.reportService.InsertUpdate(nim, input)
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	c.JSON(200, helper.SuccessResponse("succes input"))
}
