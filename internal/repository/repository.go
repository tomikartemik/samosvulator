package repository

import (
	"gorm.io/gorm"
	"samosvulator/internal/model"
)

type Repository struct {
	User
	Record
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:   NewUserRepository(db),
		Record: NewRecordRepository(db),
	}
}

type User interface {
	SignUp(user model.User) error
	SignIn(userData model.SignInInput) (int, error)
	GetUserByID(userID int) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
	ChangePassword(userID int, password string) error
}

type Record interface {
	CreateRecord(record model.Record) error
	GetAllRecords() ([]model.Record, error)
	GetRecordsByUserID(userID int) ([]model.Record, error)
}
