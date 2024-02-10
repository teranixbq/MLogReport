package response

type ResponseLogin struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type ProfileUser struct {
	Nim         string  `json:"nim"`
	Name        string  `json:"name"`
	Class       string  `json:"class"`
	Program     string  `json:"program"`
	Total_Score float64 `json:"total_score"`
}
