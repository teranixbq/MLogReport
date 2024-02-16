package handler

import (
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

	fileTranscript, _ := c.FormFile("transcript")
	fileFinalReport, _ := c.FormFile("final_report")
	fileCertification, _ := c.FormFile("certification")

	result, err := report.reportService.InsertUpdate(nim, fileFinalReport, fileTranscript, fileCertification)
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			c.AbortWithStatusJSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.AbortWithStatusJSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	c.JSON(200, helper.SuccessWithDataResponse("succes input file", result))
}
