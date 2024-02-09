package request

import "mlogreport/feature/admin/model"

func CreateAdvisorToModel(data CreateAdvisor) model.Admin {
	return model.Admin{
		Nip:      data.Nip,
		Name:     data.Name,
		Password: data.Password,
		Role:     data.Role,
	}
}
