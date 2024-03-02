package response

import "time"

type ResponseWeekly struct {
	Id          string    `json:"id"`
	Name        string    `json:"name,omitempty"`
	UsersId     string    `json:"users_id,omitempty"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type ResponseWeeklyDetail struct {
	UsersId string           `json:"users_id"`
	Name    string           `json:"name"`
	Data    []ResponseWeekly `json:"data_weekly"`
}

type ResponsePeriode struct {
	Id        string    `json:"id"`
	Start     string    `json:"start"`
	End       string    `json:"end"`
}
