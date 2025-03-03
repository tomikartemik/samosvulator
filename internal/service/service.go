package service

import (
	"samosvulator/internal/model"
	"samosvulator/internal/repository"
)

type Service struct {
	User
	Record
	Sheets
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:   NewUserService(repos.User),
		Record: NewRecordService(repos.Record),
		Sheets: NewSheetsService(repos.Record, repos.User),
	}
}

type User interface {
	SignUp(userData model.User) error
	SignIn(userData model.SignInInput) (model.UserOutput, error)
}

type Record interface {
	CreateRecord(record model.Record) error
	GetAllRecords() ([]model.Record, error)
}

type Sheets interface {
	GetRecordsForAnalise() ([]model.RecordForAnalise, error)
}
