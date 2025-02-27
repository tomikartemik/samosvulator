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
		Records:  user.Records, // Копируем записи, если они есть
	}
}
