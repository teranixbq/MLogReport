package service

import (
	"mlogreport/feature/report/dto/request"
	"mlogreport/feature/report/dto/response"
	"mlogreport/feature/report/repository"
)

type reportService struct {
	reportRepository repository.ReportRepositoryInterface
}

type ReportServiceInterface interface {
	InsertUpdate(nim string, filepdf request.RequestReportFile) (response.ResponseReport, error)
}

func NewReportService(reportRepository repository.ReportRepositoryInterface) ReportServiceInterface {
	return &reportService{
		reportRepository: reportRepository,
	}
}

func (report *reportService) InsertUpdate(nim string, filepdf request.RequestReportFile) (response.ResponseReport, error) {

	filepdf.FinalReport.Filename = "FR-" + nim
	filepdf.Transcript.Filename = "TR-" + nim
	filepdf.Certification.Filename = "CR-" + nim

	dataReport, err := report.reportRepository.InsertUpdate(nim, filepdf)
	if err != nil {
		return response.ResponseReport{}, nil
	}

	return dataReport, nil
}
