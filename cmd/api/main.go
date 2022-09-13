package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New()

	server.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON("pong :)")
	})

	if err := server.Listen(":3001"); err != nil {
		log.Fatalf("Server was not able to start. Error: %v", err.Error())
	}
}
