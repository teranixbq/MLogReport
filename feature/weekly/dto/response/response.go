package response

import "time"

type ResponseWeekly struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	UsersId     string    `json:"users_id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
