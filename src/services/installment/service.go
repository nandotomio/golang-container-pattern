package installment

import (
	"errors"
	"time"

	"github.com/nandotomio/golang-container-pattern/src/interfaces"
	"github.com/nandotomio/golang-container-pattern/src/repositories"
	"github.com/nandotomio/golang-container-pattern/src/structs"
)

type InstallmentService struct {
	InstallmentRepository interfaces.InstallmentRepository
}

func NewInstallmentService(repositoryContainer repositories.RepositoryContainer) InstallmentService {
	return InstallmentService{
		InstallmentRepository: repositoryContainer.InstallmentRepository,
	}
}

func (is InstallmentService) Create(installment structs.Installment) error {
	if installment.DueDay < time.Now().Day() {
		return errors.New("Installment expired")
	}

	return is.InstallmentRepository.Create(installment)
}
