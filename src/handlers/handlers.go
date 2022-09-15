package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nandotomio/golang-container-pattern/src/handlers/installment"
	"github.com/nandotomio/golang-container-pattern/src/services"
)

func NewHandlerContainer(router fiber.Router, serviceContainer services.ServiceContainer) {
	installment.NewInstallmentHandler(router, serviceContainer).SetRoutes()
}
