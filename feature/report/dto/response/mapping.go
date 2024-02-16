package response

import "mlogreport/feature/report/model"

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
