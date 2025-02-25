package usecases

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
)

type BuscarClienteUseCase struct {
	clienteRepository repository.ClienteRepository
}

func NewBuscarClienteUseCase(clienteRepository repository.ClienteRepository) *BuscarClienteUseCase {
	return &BuscarClienteUseCase{
		clienteRepository: clienteRepository,
	}
}

func (uc *BuscarClienteUseCase) BuscarCliente(c context.Context, cpf string) (entities.Cliente, error) {
	return uc.clienteRepository.BuscarCliente(c, cpf)
}