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
	InsertUpdate(nim string, filepdf request.RequestReportFile) (response.ResponseReport, error)
}

func NewReportRepository(db *gorm.DB, sb storage.StorageInterface) ReportRepositoryInterface {
	return &reportRepository{
		db: db,
		sb: sb,
	}
}

func (report *reportRepository) InsertUpdate(nim string, filepdf request.RequestReportFile) (response.ResponseReport, error) {
	data := request.RequestReport{}
	request := request.RequestReportToModel(nim, data)

	fileNames := map[int]string{
		0: "FR",
		1: "TR",
		2: "CR",
	}

	fileinput := []*multipart.FileHeader{
		filepdf.FinalReport,
		filepdf.Transcript,
		filepdf.Certification,
	}

	for i, file := range fileinput {
		if file != nil {
			url, err := report.sb.Upload(file)
			if err != nil {
				return response.ResponseReport{}, err
			}
			switch fileNames[i] {
			case "FR":
				request.FinalReport = url
			case "TR":
				request.Transcript = url
			case "CR":
				request.Certification = url
			}
		}
	}

	tx := report.db.Create(&request)
	if tx.Error != nil {
		return response.ResponseReport{}, tx.Error
	}

	response := response.ModelToResponseReport(request)
	return response, nil
}
