package usecases

import (
	"context"
	"lanchonete/domain"
	"lanchonete/domain/repository"
	"lanchonete/domain/entities"
	"lanchonete/usecases/cliente"
)

type ClienteUseCaseImpl struct {
	criarClienteUC  *usecases.CriarClienteUseCase
	buscarClienteUC *usecases.BuscarClienteUseCase
}

func NewClienteUseCase(clienteRepository repository.ClienteRepository) domain.ClienteUseCase {
	return &ClienteUseCaseImpl{
		criarClienteUC:  usecases.NewCriarClienteUseCase(clienteRepository),
		buscarClienteUC: usecases.NewBuscarClienteUseCase(clienteRepository),
	}
}

func (uc *ClienteUseCaseImpl) CriarCliente(c context.Context, cliente *entities.Cliente) error {
	return uc.criarClienteUC.CriarCliente(c, cliente)
}

func (uc *ClienteUseCaseImpl) BuscarCliente(c context.Context, cpf string) (entities.Cliente, error) {
	return uc.buscarClienteUC.BuscarCliente(c, cpf)
}