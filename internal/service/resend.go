package service

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"net/smtp"
	"os"
	"samosvulator/internal/repository"
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
	smtpPort := "587"
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	// MIME-заголовки для HTML-письма
	header := fmt.Sprintf(
		"To: %s\r\n"+
			"From: %s\r\n"+
			"Subject: Ваш новый пароль\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=\"UTF-8\"\r\n"+
			"\r\n", mail, smtpUser)

	// HTML-тело письма
	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<style>
				body { font-family: Arial, sans-serif; color: #333; }
				.container { max-width: 600px; margin: 0 auto; padding: 20px; background-color: #f9f9f9; border-radius: 10px; }
				.header { background-color: #4285f4; color: white; padding: 15px; text-align: center; border-radius: 10px 10px 0 0; }
				.content { padding: 20px; background-color: white; border-radius: 0 0 10px 10px; }
				.password { font-size: 18px; font-weight: bold; color: #4285f4; background-color: #e8f0fe; padding: 10px; border-radius: 5px; text-align: center; }
				.footer { text-align: center; font-size: 12px; color: #777; margin-top: 20px; }
				.button { display: inline-block; padding: 10px 20px; background-color: #4285f4; color: white; text-decoration: none; border-radius: 5px; }
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<h2>Восстановление пароля</h2>
				</div>
				<div class="content">
					<p>Здравствуйте!</p>
					<p>Вы запросили восстановление пароля. Вот ваш новый пароль:</p>
					<div class="password">%s</div>
					<p>Используйте его для входа в систему. Рекомендуем сменить пароль после входа для повышения безопасности.</p>
				</div>
				<div class="footer">
					<p>Если вы не запрашивали восстановление, проигнорируйте это письмо или свяжитесь с поддержкой.</p>
					<p>&copy; 2025 Samosvulator</p>
				</div>
			</div>
		</body>
		</html>
	`, newPassword)

	msg := []byte(header + body)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, []string{mail}, msg)
	if err != nil {
		return err
	}

	fmt.Println("service resend " + newPassword)
	//hashedPassword := utils.GeneratePasswordHash(newPassword)
	//fmt.Println("service resend hashed " + hashedPassword)
	err = s.repo.ChangePassword(user.ID, newPassword)

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

	// Заполняем первые 3 символа из категорий
	categories := []string{upper, lower, digits}
	for i, cat := range categories {
		char, err := randomChar(cat)
		if err != nil {
			return "", err
		}
		password[i] = char
	}

	// Заполняем оставшиеся 13 символов (с 3 до 15)
	for i := 3; i < 16; i++ { // Изменено с 4 на 3, чтобы заполнить все 16 байтов
		char, err := randomChar(all)
		if err != nil {
			return "", err
		}
		password[i] = char
	}

	// Перемешиваем
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
