package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	server := fiber.New()

	postgresDns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		"dev",
		"change_please",
		"dev",
		"5432",
	)
	postgresDialector := postgres.Open(postgresDns)
	_, err := gorm.Open(postgresDialector, &gorm.Config{})

	if err != nil {
		log.Fatalf("DB connection error. Error: %v", err.Error())
	}

	server.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON("pong :)")
	})

	if err := server.Listen(":3001"); err != nil {
		log.Fatalf("Server was not able to start. Error: %v", err.Error())
	}
}
