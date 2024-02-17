package service

import (
	"mime/multipart"
	"mlogreport/feature/report/dto/response"
	"mlogreport/feature/report/repository"
)

type reportService struct {
	reportRepository repository.ReportRepositoryInterface
}

type ReportServiceInterface interface {
	InsertUpdate(nim string, finalReport, transcript, certification *multipart.FileHeader) (response.ResponseReport, error)
}

func NewReportService(reportRepository repository.ReportRepositoryInterface) ReportServiceInterface {
	return &reportService{
		reportRepository: reportRepository,
	}
}

func (report *reportService) InsertUpdate(nim string, finalReport, transcript, certification *multipart.FileHeader) (response.ResponseReport, error) {
	finalReport.Filename = "FR-"+nim
	transcript.Filename = "TR-"+nim
	certification.Filename = "CR-"+nim

	dataReport, err := report.reportRepository.InsertUpdate(nim, finalReport, transcript, certification)
	if err != nil {
		return response.ResponseReport{}, nil
	}

	return dataReport, nil
}
