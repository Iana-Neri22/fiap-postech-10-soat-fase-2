package usecases

import (
	"context"

	"lanchonete/domain/repository"
	"lanchonete/domain/usecases"
	"lanchonete/domain/entities"
)

type clienteUseCase struct {
	clienteRepository repository.ClienteRepository
}

func NewClienteUseCase(clienteRepository repository.ClienteRepository) usecases.ClienteUseCase {
	return &clienteUseCase{
		clienteRepository: clienteRepository,
	}
}

func (u *clienteUseCase) CriarCliente(c context.Context, cliente *entities.Cliente) error {
	return u.clienteRepository.CriarCliente(c, cliente)
}

func (u *clienteUseCase) BuscarCliente(c context.Context, CPF string) (entities.Cliente, error) {
	return u.clienteRepository.BuscarCliente(c, CPF)
}