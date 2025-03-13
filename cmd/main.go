package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/resend/resend-go/v2"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"samosvulator/internal"
	"samosvulator/internal/handler"
	"samosvulator/internal/repository"
	"samosvulator/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables.")
	}

	apiKey := os.Getenv("RESEND_API_KEY")

	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "Acme <onboarding@resend.dev>",
		To:      []string{"artemules@mail.ru"},
		Subject: "Hello world",
		Html:    "<strong>It works!</strong>",
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		panic(err)
	}
	fmt.Println(sent.Id)

	db, err := repository.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Application started successfully")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(internal.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running server %s", err.Error())
	}
}
