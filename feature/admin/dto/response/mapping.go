package response

import (
	"mlogreport/feature/admin/model"
	user "mlogreport/feature/user/dto/response"
)

func ModelToResponseLogin(name, roles, token string) ResponseLogin {
	return ResponseLogin{
		Name:  name,
		Roles: roles,
		Token: token,
	}
}

func ModelToResponseAllAdvisor(data model.Admins) ResponseAllAdvisor {
	return ResponseAllAdvisor{
		Id:   data.Id,
		Nip:  data.Nip,
		Name: data.Name,
	}
}

func ListResponseAllAdvisor(data []model.Admins) []ResponseAllAdvisor {
	list := []ResponseAllAdvisor{}
	for _, v := range data {
		response := ModelToResponseAllAdvisor(v)
		list = append(list, response)
	}

	return list
}

func ModelToResponseAdvisor(data model.Admins) ResponseAdvisor {
	return ResponseAdvisor{
		Id:       data.Id,
		Nip:      data.Nip,
		Name:     data.Name,
		Colleges: user.ListModelToProfileUser(data.Advisor),
	}
}
