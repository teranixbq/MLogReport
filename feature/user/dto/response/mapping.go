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
		Id:     data.Id,
		Name:    data.Name,
		Class:   data.Class,
		Mitra:   data.Mitra,
		Program: data.Program,
	}
}
