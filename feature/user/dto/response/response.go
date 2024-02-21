package response

type ResponseLogin struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type ProfileUser struct {
	Nim     string `json:"nim"`
	Name    string `json:"name"`
	Class   string `json:"class"`
	Mitra   string `json:"mitra"`
	Program string `json:"program"`
}
