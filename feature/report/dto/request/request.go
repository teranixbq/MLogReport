package request

import "mime/multipart"

type RequestReport struct {
	FinalReport   string `form:"final_report"`
	Transcript    string `form:"transcript"`
	Certification string `form:"certification"`
}

type RequestReportFile struct {
	FinalReport   *multipart.FileHeader 
	Transcript    *multipart.FileHeader 
	Certification *multipart.FileHeader 
}
