package request

import (
	"mime/multipart"
	"mlogreport/feature/report/model"
)

func RequestReportToModel(userid string, data RequestReport) model.Report {
	return model.Report{
		UsersId:       userid,
		FinalReport:   data.FinalReport,
		Transcript:    data.Transcript,
		Certification: data.Certification,
	}
}

func MultipartToRequestReport(finalreport, transcript, certification *multipart.FileHeader) RequestReportFile {
	return RequestReportFile{
		FinalReport:   finalreport,
		Transcript:    transcript,
		Certification: certification,
	}
}
