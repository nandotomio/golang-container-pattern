package services

import (
	"github.com/nandotomio/golang-container-pattern/src/interfaces"
	"github.com/nandotomio/golang-container-pattern/src/repositories"
)

type ServiceContainer struct {
	InstallmentService interfaces.InstallmentService
}

func GetService(repositoryContainer repositories.RepositoryContainer) ServiceContainer {
	return ServiceContainer{}
}
