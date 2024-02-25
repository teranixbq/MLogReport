package response

import "mlogreport/feature/admin/model"

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
