package repository

import (
	"mime/multipart"
	"mlogreport/app/storage"
	admin "mlogreport/feature/admin/model"
	"mlogreport/feature/report/dto/request"
	"mlogreport/feature/report/dto/response"
	"mlogreport/feature/report/model"
	user "mlogreport/feature/user/model"

	"gorm.io/gorm"
)

type reportRepository struct {
	db *gorm.DB
	sb storage.StorageInterface
}

type ReportRepositoryInterface interface {
	InsertUpdate(nim string, filepdf request.RequestReportFile) error
	FindReport(nim string) (response.ResponseReport, error)
	FindAllReport(nip string) ([]response.ResponseReport, error)
}

func NewReportRepository(db *gorm.DB, sb storage.StorageInterface) ReportRepositoryInterface {
	return &reportRepository{
		db: db,
		sb: sb,
	}
}

func (report *reportRepository) InsertUpdate(IdUser string, filepdf request.RequestReportFile) error {
	data := request.RequestReport{}
	request := request.RequestReportToModel(IdUser, data)

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
				return err
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

	dataReport, err := report.FindReport(IdUser)
	if err != nil {
		tx := report.db.Create(&request)
		if tx.Error != nil {
			return tx.Error
		}
	}

	tx := report.db.Where("id= ?", dataReport.Id).Updates(&request)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (report *reportRepository) FindReport(IdUser string) (response.ResponseReport, error) {
	dataReport := model.Report{}

	tx := report.db.Where("users_id = ? ", IdUser).Take(&dataReport)
	if tx.Error != nil {
		return response.ResponseReport{}, tx.Error
	}

	response := response.ModelToResponseReport(dataReport)
	return response, nil
}

func (report *reportRepository) FindAllReport(IdUser string) ([]response.ResponseReport, error) {
	dataReport := []response.ResponseReport{}
	dataAdmin := admin.Admins{}
	dataUser := []user.Users{}

	var unsubmitted int

	tx := report.db.Where("id = ?", IdUser).First(&dataAdmin)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if txUser := report.db.Model(&dataAdmin).Association("Advisor").Find(&dataUser); txUser != nil {
		return nil, txUser
	}

	for _, user := range dataUser {
		var userReport model.Report

		err := report.db.Where("users_id = ?", user.Id).First(&userReport)

		if err.RowsAffected == 0 {
			unsubmitted++
			continue
		}
		if err.Error != nil {
			continue
		}

		result := response.ModelToResponseReportUser(userReport)
		result.Name = user.Name
		dataReport = append(dataReport, result)
	}

	return dataReport, nil
}
