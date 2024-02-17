package service

import (
	"mime/multipart"
	"mlogreport/feature/report/dto/request"
	"mlogreport/feature/report/repository"
)

type reportService struct {
	reportRepository repository.ReportRepositoryInterface
}

type ReportServiceInterface interface {
	InsertUpdate(nim string, filepdf request.RequestReportFile) error
}

func NewReportService(reportRepository repository.ReportRepositoryInterface) ReportServiceInterface {
	return &reportService{
		reportRepository: reportRepository,
	}
}

func (report *reportService) InsertUpdate(nim string, filepdf request.RequestReportFile) error {

	fileinput := []*multipart.FileHeader{
		filepdf.FinalReport,
		filepdf.Transcript,
		filepdf.Certification,
	}

	for i, file := range fileinput {
		if file != nil {
			switch i {
			case 0:
				file.Filename = "FR-" + nim
			case 1:
				file.Filename = "TR-" + nim
			case 2:
				file.Filename = "CR-" + nim
			}
		}
	}

	err := report.reportRepository.InsertUpdate(nim, filepdf)
	if err != nil {
		return err
	}

	return nil
}
