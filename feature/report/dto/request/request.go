package request

type RequestReport struct {
	FinalReport   string `form:"final_report"`
	Transcript    string `form:"transcript"`
	Certification string `form:"certification"`
}
