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

// func (report *reportRepository) InsertUpdate(nim string, file *multipart.FileHeader, data request.RequestReport) (response.ResponseReport, error) {
// 	request := request.RequestReportToModel(nim, data)

// 	url, err := report.sb.Upload(file)

// 	tx := report.db.Create(&request)
// 	if tx.Error != nil {
// 		return response.ResponseReport{}, tx.Error
// 	}

// 	response := response.ModelToResponseReport(request)
// 	return response,nil
// }

func (report *reportRepository) InsertUpdate(nim string, fileFinalReport, fileTranscript, fileCertification *multipart.FileHeader) (response.ResponseReport, error) {
    data := request.RequestReport{}
	request := request.RequestReportToModel(nim, data)

    // Skip jika fileFinalReport nil
    if fileFinalReport == nil {
        request.FinalReport = ""
    } else {
        urlFinalReport, err := report.sb.Upload(fileFinalReport)
        if err != nil {
            return response.ResponseReport{}, err
        }
        request.FinalReport = urlFinalReport
    }

    // Skip jika fileTranscript nil
    if fileTranscript == nil {
        request.Transcript = ""
    } else {
        urlTranscript, err := report.sb.Upload(fileTranscript)
        if err != nil {
            return response.ResponseReport{}, err
        }
        request.Transcript = urlTranscript
    }

    // Skip jika fileCertification nil
    if fileCertification == nil {
        request.Certification = ""
    } else {
        urlCertification, err := report.sb.Upload(fileCertification)
        if err != nil {
            return response.ResponseReport{}, err
        }
        request.Certification = urlCertification
    }

    // Simpan model Report ke database
    tx := report.db.Create(&request)
    if tx.Error != nil {
        return response.ResponseReport{}, tx.Error
    }

    response := response.ModelToResponseReport(request)
    return response, nil
}
