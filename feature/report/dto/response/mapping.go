package response

import (
	"mlogreport/feature/report/model"
)

func ModelToResponseReport(data model.Report) ResponseReport {
	return ResponseReport{
		Id:            data.Id,
		FinalReport:   data.FinalReport,
		Transcript:    data.Transcript,
		Certification: data.Certification,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}
}

func ModelToResponseReportUser(data model.Report) ResponseReport {
	return ResponseReport{
		Id:            data.Id,
		Name:          "",
		UsersId:       data.UsersId,
		FinalReport:   data.FinalReport,
		Transcript:    data.Transcript,
		Certification: data.Certification,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}
}

func ListModelToResponseReport(data []model.Report) []ResponseReport {
	list := []ResponseReport{}
	for _, v := range data {
		result := ModelToResponseReportUser(v)
		list = append(list, result)
	}
	return list
}
