package response

import "mlogreport/feature/user/model"

func ModelToResponseLogin(data model.Users, token string) ResponseLogin {
	return ResponseLogin{
		Name:  data.Name,
		Token: token,
	}
}

func ModelToProfileUser(data model.Users) ProfileUser {
	return ProfileUser{
		Nim:     data.Nim,
		Name:    data.Name,
		Class:   data.Class,
		Mitra:   data.Mitra,
		Program: data.Program,
	}
}

func ListModelToProfileUser(data []model.Users) []ProfileUser {
	list := []ProfileUser{}
	for _, v := range data {
		response := ModelToProfileUser(v)
		list = append(list, response)
	}
	return list
}
