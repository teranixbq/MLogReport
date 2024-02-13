package response

import "mlogreport/feature/weekly/model"

func ModelToResponseWeekly(data model.Weekly) ResponseWeekly {
	return ResponseWeekly{
		Id:          data.Id,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
	}
}

func ListModelToResponseWeekly (data []model.Weekly) []ResponseWeekly{
	listweekly := []ResponseWeekly{}
	for _,v := range data {
		response := ModelToResponseWeekly(v)
		listweekly = append(listweekly, response)
	}
	return listweekly
}