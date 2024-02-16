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

	// fileTranscript, err := c.FormFile("transcript")
	// if err != nil {
	// 	c.AbortWithStatusJSON(400, helper.ErrorResponse("2"))
	// 	return
	// }

	// fileFinalReport, err := c.FormFile("final_report")
	// if err != nil {
	// 	c.AbortWithStatusJSON(400, helper.ErrorResponse("1"))
	// 	return
	// }

	fileCertification, _ := c.FormFile("certification")

	result, err := report.reportService.InsertUpdate(nim, nil, nil, fileCertification)
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
