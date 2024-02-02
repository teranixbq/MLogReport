package request

type CreateAdvisor struct {
	Nip      string `json:"nip"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AdminLogin struct {
	Nip             string `json:"nip"`
	Password        string `json:"password"`
}
