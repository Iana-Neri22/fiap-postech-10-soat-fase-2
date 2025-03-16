package usecase

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	"lanchonete/domain/usecases"
)

type clienteUseCase struct {
	clienteRepository repository.ClienteRepository
}

func NewClienteUseCase(clienteRepository repository.ClienteRepository) usecases.ClienteUseCase {
	return &clienteUseCase{
		clienteRepository: clienteRepository,
	}
}

func (cuc *clienteUseCase) CriarCliente(c context.Context, cliente *entities.Cliente) error {
	return cuc.clienteRepository.CriarCliente(c, cliente)
}

func (cuc *clienteUseCase) BuscarCliente(c context.Context, CPF string) (entities.Cliente, error) {
	return cuc.clienteRepository.BuscarCliente(c, CPF)
}

