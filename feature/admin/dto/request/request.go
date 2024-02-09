package request

type CreateAdvisor struct {
	Nip      string `json:"nip"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type AdminLogin struct {
	Nip      string `json:"nip"`
	Password string `json:"password"`
}

type ListCollege struct {
	Advisor string `json:"advisor"`
	Colleges []string `json:"colleges"`
}
