package service

import (
	"fmt"
	"samosvulator/internal/model"
	"samosvulator/internal/repository"
	"samosvulator/internal/utils"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SignUp(userData model.User) error {
	userData.Password = utils.GeneratePasswordHash(userData.Password)
	return s.repo.SignUp(userData)
}

func (s *UserService) SignIn(userData model.SignInInput) (model.SignInOutput, error) {
	userID, err := s.repo.SignIn(userData)
	if err != nil {
		return model.SignInOutput{}, err
	}

	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return model.SignInOutput{}, err
	}

	token, err := CreateToken(user.ID)
	if err != nil {
		return model.SignInOutput{}, err
	}

	userOutput := model.SignInOutput{
		Token: token,
		User:  utils.UserToUserOutput(user),
	}

	return userOutput, nil
}
