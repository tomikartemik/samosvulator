package repository

import (
	"errors"
	"gorm.io/gorm"
	"samosvulator/internal/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) SignUp(user model.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) SignIn(userData model.SignInInput) (int, error) {
	var user model.User

	if err := r.db.Where("username = ?", userData.Username).First(&user).Error; err != nil {
		return 0, errors.New("Пользователя с таким никнеймом не существует!")
	}

	if user.Password != userData.Password {
		return 0, errors.New("Неверный пароль!")
	}

	return user.ID, nil
}

func (r *UserRepository) GetUserByID(userID int) (model.User, error) {
	var user model.User
	if err := r.db.Where("id = ?", userID).Preload("Records").First(&user).Error; err != nil {
		return model.User{}, errors.New("Пользователь с таким ID не найден!")
	}
	return user, nil
}
