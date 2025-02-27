package service

import (
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

func (s *UserService) SignIn(userData model.SignInInput) (model.UserOutput, error) {
	userData.Password = utils.GeneratePasswordHash(userData.Password)

	userID, err := s.repo.SignIn(userData)
	if err != nil {
		return model.UserOutput{}, err
	}

	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return model.UserOutput{}, err
	}

	return utils.UserToUserOutput(user), nil
}
