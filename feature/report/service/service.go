package service

import (
	"errors"
	"mime/multipart"
	"mlogreport/feature/report/dto/request"
	"mlogreport/feature/report/dto/response"
	"mlogreport/feature/report/repository"
	"mlogreport/utils/validation"
)

type reportService struct {
	reportRepository repository.ReportRepositoryInterface
}

type ReportServiceInterface interface {
	InsertUpdate(nim string, filepdf request.RequestReportFile) error
	FindReport(nim string) (response.ResponseReport, error)
	FindAllReport(nip string) ([]response.ResponseReport, error)
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
	errEmpty := validation.CheckAllEmpty(filepdf)
	if errEmpty != nil {
		return errors.New("error : all data cannot be empty, there must be at least 1")
	}

	for _, v := range fileinput {
		if v != nil {
			if v.Size > 10*1024*1024 {
				return errors.New("error : file size cannot be more than 10MB")
			}

			contentType := v.Header.Get("Content-Type")
			if contentType != "application/pdf" {
				return errors.New("error : file type must be pdf")
			}
		}
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

func (report *reportService) FindReport(nim string) (response.ResponseReport, error) {
	data, err := report.reportRepository.FindReport(nim)
	if err != nil {
		return response.ResponseReport{}, err
	}

	return data, nil
}

func (report *reportService) FindAllReport(nip string) ([]response.ResponseReport, error) {
	dataReport,err := report.reportRepository.FindAllReport(nip)
	if err != nil {
		return nil,nil
	}

	return dataReport,nil
}