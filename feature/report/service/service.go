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
	InsertUpdate(nim string, fileFinalReport, fileTranscript, fileCertification *multipart.FileHeader) (response.ResponseReport, error)
}

func NewReportService(reportRepository repository.ReportRepositoryInterface) ReportServiceInterface {
	return &reportService{
		reportRepository: reportRepository,
	}
}

func (report *reportService) InsertUpdate(nim string, fileFinalReport *multipart.FileHeader, fileTranscript *multipart.FileHeader, fileCertification *multipart.FileHeader) (response.ResponseReport, error) {

	dataReport,err :=report.reportRepository.InsertUpdate(nim,fileFinalReport,fileTranscript,fileCertification) 
	if err != nil {
		return response.ResponseReport{},nil
	}

	return dataReport, nil
}