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
		Nim:         data.Nim,
		Name:        data.Name,
		Class:       data.Class,
		Program:     data.Program,
		Total_Score: data.Total_Score,
	}
}
