package response

import "mlogreport/feature/user/model"

func ModelToResponseLogin(data model.Users,token string) ResponseLogin{
	return ResponseLogin{
		Name:  data.Name,
		Token: token,
	}
}