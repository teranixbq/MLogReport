package meta

import "math"

type Meta struct {
	Current   int `json:"current_page"`
	Limit     int `json:"limit"`
	TotalData int `json:"total_data"`
	TotalPage int `json:"total_page"`
}

func MetaInfo(page, limit, totalData int) Meta {
	totalPages := int(math.Ceil(float64(totalData) / float64(limit)))
	return Meta{
		Current:   page,
		Limit:     limit,
		TotalData: totalData,
		TotalPage: totalPages,
	}
}
