package response

import "mlogreport/feature/weekly/model"

func ModelToResponseWeekly(data model.Weekly) ResponseWeekly {
	return ResponseWeekly{
		Id:          data.Id,
		Description: data.Description,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt,
	}
}

func ModelToResponseWeeklyAdvisor(data model.Weekly) ResponseWeekly {
	return ResponseWeekly{
		Id:          data.Id,
		Description: data.Description,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt,
	}
}

func ModelToResponseWeeklyDetail(nim, name string, data []model.Weekly) ResponseWeeklyDetail {
	return ResponseWeeklyDetail{
		UsersId: nim,
		Name:    name,
		Data:    ListModelToResponseWeekly(data),
	}
}

func ListModelToResponseWeekly(data []model.Weekly) []ResponseWeekly {
	listweekly := []ResponseWeekly{}
	for _, v := range data {
		response := ModelToResponseWeeklyAdvisor(v)
		listweekly = append(listweekly, response)
	}
	return listweekly
}
