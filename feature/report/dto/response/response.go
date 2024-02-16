package response

import "time"

type ResponseReport struct {
	Id            string    `json:"id"`
	FinalReport   string    `json:"final_report"`
	Transcript    string    `json:"transcript"`
	Certification string    `json:"certification"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
