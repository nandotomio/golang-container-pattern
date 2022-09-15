package installment

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nandotomio/golang-container-pattern/src/interfaces"
	"github.com/nandotomio/golang-container-pattern/src/services"
	"github.com/nandotomio/golang-container-pattern/src/structs"
)

type InstallmentHandler struct {
	router             fiber.Router
	installmentService interfaces.InstallmentService
}

func NewInstallmentHandler(router fiber.Router, serviceContainer services.ServiceContainer) InstallmentHandler {
	return InstallmentHandler{
		router:             router,
		installmentService: serviceContainer.InstallmentService,
	}
}

func (ih InstallmentHandler) SetRoutes() {
	group := ih.router.Group("/installment")
	group.Post("", ih.CreateInstallment)
}

func (ih InstallmentHandler) CreateInstallment(c *fiber.Ctx) error {
	var body structs.Installment

	err := c.BodyParser(&body)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	err = ih.installmentService.Create(body)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON("Installment created")
}
