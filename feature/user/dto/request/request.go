package request

type RequestUser struct {
	Nim         string  `json:"nim"`
	Name        string  `json:"name"`
	Password    string  `json:"password"`
	Class       string  `json:"class"`
}

type RequestLogin struct {
	Nim      string `json:"nim"`
	Password string `json:"password"`
}