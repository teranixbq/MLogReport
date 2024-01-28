package request

type RequestUser struct {
	Nim         string  `json:"nim"`
	Name        string  `json:"name"`
	Password    string  `json:"password"`
	Class       string  `json:"class"`
	Program     string  `json:"program"`
	Total_Score float64 `json:"total_score"`
}