package request

type RequestWeekly struct {
	UsersId     string `json:"users_id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
