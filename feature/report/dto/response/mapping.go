package response

import (
	"mlogreport/feature/report/model"
	"mlogreport/utils/constanta"
)

func ModelToResponseReport(data model.Report) ResponseReport {
	url := constanta.URL_STORAGE
	return ResponseReport{
		Id:            data.Id,
		FinalReport:   url+data.FinalReport,
		Transcript:    url+data.Transcript,
		Certification: url+data.Certification,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}
}
