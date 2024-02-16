package request

import "mlogreport/feature/report/model"

func RequestReportToModel(userid string, data RequestReport) model.Report {
	return model.Report{
		UsersId:       userid,
		FinalReport:   data.FinalReport,
		Transcript:    data.Transcript,
		Certification: data.Certification,
	}
}
