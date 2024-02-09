package request

import (
	"mlogreport/feature/admin/model"
	//user "mlogreport/feature/user/model"
)

func CreateAdvisorToModel(data CreateAdvisor) model.Admin {
	return model.Admin{
		Nip:      data.Nip,
		Name:     data.Name,
		Password: data.Password,
		Role:     data.Role,
	}
}

// func RequestListCollegeToModel(data ListCollege) model.AdvisorCollege{
// 	return model.AdvisorCollege{
// 		AdminNip: data.Nip,
// 		UsersNim: data.College,
// 	}
// }

// func UserToListCollege(data Advisorcollege) model.AdvisorCollege {
// 	return model.AdvisorCollege{
// 		UsersNim: data.Nim,
// 	}
// }

// func ListColleges(data []Advisorcollege) []model.AdvisorCollege{
// 	list := []model.AdvisorCollege{}
// 	for _,v := range data {
// 		result := UserToListCollege(v)
// 		list = append(list, result)
// 	}

// 	return list
// }