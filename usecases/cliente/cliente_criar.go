package usecases

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
)

type CriarClienteUseCase struct {
	clienteRepository repository.ClienteRepository
}

func NewCriarClienteUseCase(clienteRepository repository.ClienteRepository) *CriarClienteUseCase {
	return &CriarClienteUseCase{
		clienteRepository: clienteRepository,
	}
}

func (uc *CriarClienteUseCase) CriarCliente(c context.Context, cliente *entities.Cliente) error {
	return uc.clienteRepository.CriarCliente(c, cliente)
}