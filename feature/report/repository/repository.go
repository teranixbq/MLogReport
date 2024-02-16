package repository

import (
	"mime/multipart"
	"mlogreport/app/storage"
	"mlogreport/feature/report/dto/request"
	"mlogreport/feature/report/dto/response"

	"gorm.io/gorm"
)

type reportRepository struct {
	db *gorm.DB
	sb storage.StorageInterface
}

type ReportRepositoryInterface interface {
	InsertUpdate(nim string, fileFinalReport, fileTranscript, fileCertification *multipart.FileHeader) (response.ResponseReport, error)
}

func NewReportRepository(db *gorm.DB, sb storage.StorageInterface) ReportRepositoryInterface {
	return &reportRepository{
		db: db,
		sb: sb,
	}
}

func (report *reportRepository) InsertUpdate(nim string, fileFinalReport, fileTranscript, fileCertification *multipart.FileHeader) (response.ResponseReport, error) {
	data := request.RequestReport{}
	request := request.RequestReportToModel(nim, data)

	var urls []string

	for _, file := range []*multipart.FileHeader{fileFinalReport, fileTranscript, fileCertification} {
		if file != nil {
			url, err := report.sb.Upload(file)
			if err != nil {
				return response.ResponseReport{}, err
			}
			urls = append(urls, url)
		}
	}

	for i, url := range urls {
		switch i {
		case 0:
			request.FinalReport = url
		case 1:
			request.Transcript = url
		case 2:
			request.Certification = url
		}
	}

	tx := report.db.Create(&request)
	if tx.Error != nil {
		return response.ResponseReport{}, tx.Error
	}

	response := response.ModelToResponseReport(request)
	return response, nil
}
