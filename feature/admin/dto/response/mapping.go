package response

func ModelToResponseLogin(name, roles, token string) ResponseLogin {
	return ResponseLogin{
		Name:  name,
		Roles: roles,
		Token: token,
	}
}