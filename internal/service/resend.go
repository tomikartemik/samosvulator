package service

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/resend/resend-go/v2"
	"math/big"
	"os"
	"samosvulator/internal/repository"
	"samosvulator/internal/utils"
)

type ResendService struct {
	repo repository.User
}

func NewResendService(repo repository.User) *ResendService {
	return &ResendService{repo: repo}
}

func (s *ResendService) ChangePassword(mail string) error {
	user, err := s.repo.GetUserByUsername(mail)
	if err != nil {
		return errors.New("Пользователя с таким именем не существует!")
	}

	newPassword, err := generatePassword()
	if err != nil {
		fmt.Println("Ошибка при генерации пароля:", err)
		return err
	}

	apiKey := os.Getenv("RESEND_API_KEY")

	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "Samosvulator <onboarding@resend.dev>",
		To:      []string{mail},
		Subject: "New password",
		Html:    "<strong>" + newPassword + "</strong>",
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		panic(err)
	}
	fmt.Println(sent.Id)

	hashedPasword := utils.GeneratePasswordHash(newPassword)

	err = s.repo.ChangePassword(user.ID, hashedPasword)

	if err != nil {
		return err
	}

	return nil
}

func generatePassword() (string, error) {
	const (
		upper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		lower  = "abcdefghijklmnopqrstuvwxyz"
		digits = "0123456789"
		all    = upper + lower + digits
	)

	password := make([]byte, 16)

	categories := []string{upper, lower, digits}
	for i, cat := range categories {
		char, err := randomChar(cat)
		if err != nil {
			return "", err
		}
		password[i] = char
	}

	for i := 4; i < 16; i++ {
		char, err := randomChar(all)
		if err != nil {
			return "", err
		}
		password[i] = char
	}

	for i := len(password) - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return "", err
		}
		password[i], password[j.Int64()] = password[j.Int64()], password[i]
	}

	return string(password), nil
}

func randomChar(chars string) (byte, error) {
	idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
	if err != nil {
		return 0, err
	}
	return chars[idx.Int64()], nil
}
