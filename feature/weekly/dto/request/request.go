package request

type RequestWeekly struct {
	UsersId     string `json:"users_id"`
	Description string `json:"description"`
}

type RequestStatus struct {
	Status     string `json:"status"`
}