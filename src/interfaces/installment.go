package interfaces

import "github.com/nandotomio/golang-container-pattern/src/structs"

type InstallmentService interface {
	Create(installment structs.Installment) error
}

type InstallmentRepository interface {
	Create(installment structs.Installment) error
}
