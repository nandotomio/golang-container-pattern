package services

import (
	"github.com/nandotomio/golang-container-pattern/src/interfaces"
	"github.com/nandotomio/golang-container-pattern/src/repositories"
	"github.com/nandotomio/golang-container-pattern/src/services/installment"
)

type ServiceContainer struct {
	InstallmentService interfaces.InstallmentService
}

func GetServices(repositoryContainer repositories.RepositoryContainer) ServiceContainer {
	return ServiceContainer{
		InstallmentService: installment.InstallmentService(repositoryContainer),
	}
}
