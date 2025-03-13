package service

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"net/smtp"
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

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"                  // Используйте порт 587 для TLS
	smtpUser := os.Getenv("SMTP_USER") // Ваш Gmail адрес
	smtpPass := os.Getenv("SMTP_PASS")

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	msg := []byte(fmt.Sprintf(
		"To: %s\r\n"+
			"Subject: Ваш новый пароль\r\n"+
			"\r\n"+
			"Ваш новый пароль: %s\r\n"+
			"Пожалуйста, смените его после входа.\r\n", mail, newPassword))
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, []string{mail}, msg)
	if err != nil {
		return err
	}

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
