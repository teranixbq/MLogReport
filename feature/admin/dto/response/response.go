package response

type ResponseLogin struct {
	Name string `json:"name"`
	Roles string `json:"roles"`
	Token string `json:"token"`
}