package response

import "time"

type ResponseReport struct {
	Id            string    `json:"id"`
	Name          string    `json:"name,omitempty"`
	UsersId       string    `json:"nim,omitempty"`
	FinalReport   string    `json:"final_report,omitempty"`
	Transcript    string    `json:"transcript,omitempty"`
	Certification string    `json:"certification,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
