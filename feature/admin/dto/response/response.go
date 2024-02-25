package response

type ResponseLogin struct {
	Name  string `json:"name"`
	Roles string `json:"roles"`
	Token string `json:"token"`
}

type ResponseAllAdvisor struct {
	Id        string `json:"id"`
	Nip       string `json:"nip"`
	Name      string `json:"name"`
}
