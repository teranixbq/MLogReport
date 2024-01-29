package response

type ResponseLogin struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}
