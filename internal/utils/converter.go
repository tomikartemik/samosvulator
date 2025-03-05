package utils

import "samosvulator/internal/model"

func UserToUserOutput(user model.User) model.UserOutput {
	return model.UserOutput{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Surname:  user.Surname,
		Company:  user.Company,
		Section:  user.Section,
		JobTitle: user.JobTitle,
	}
}

func ConvertToRecordForAnalise(record model.Record, user model.User) model.RecordForAnalise {
	return model.RecordForAnalise{
		ID:             record.ID,
		ExcavatorName:  record.ExcavatorName,
		Date:           record.Date,
		Shift:          record.Shift,
		ShiftTime:      record.ShiftTime,
		LoadTime:       record.LoadTime,
		CycleTime:      record.CycleTime,
		ApproachTime:   record.ApproachTime,
		ActualTrucks:   record.ActualTrucks,
		Productivity:   record.Productivity,
		RequiredTrucks: record.RequiredTrucks,
		PlanVolume:     record.PlanVolume,
		ForecastVolume: record.ForecastVolume,
		Downtime:       record.Downtime,
		UserName:       user.Name,
		UserSurname:    user.Surname,
		Company:        user.Company,
		Section:        user.Section,
		JobTitle:       user.JobTitle,
	}
}
