package request

import "mlogreport/feature/weekly/model"

func RequestWeeklyToModel(data RequestWeekly) model.Weekly {
	return model.Weekly{
		UsersId:     data.UsersId,
		Description: data.Description,
	}
}

func RequestPeriodeToModel(data RequestPeriode) model.Periode {
	return model.Periode{
		Start: data.Start,
		End:   data.End,
	}
}
