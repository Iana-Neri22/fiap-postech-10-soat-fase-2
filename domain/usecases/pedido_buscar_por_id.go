package usecases

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
)

type PedidoBuscarPorIdUseCase struct {
	pedidoGateway repository.PedidoRepository
}

<<<<<<< HEAD
func (pduc *PedidoBuscarPorIdUseCase) Run(c context.Context, identificacao string) (*entities.Pedido, error) {
	pedido, err := pduc.pedidoGateway.BuscarPedido(c, identificacao)
	if err != nil {
		return nil, err
=======
func (pduc *PedidoBuscarPorIdUseCase) Run(c context.Context, identificacao string) (entities.Pedido, error) {
	pedido, err := pduc.pedidoGateway.BuscarPedido(c, identificacao)
	if err != nil {
		return entities.Pedido{}, err
>>>>>>> iana
	}
	return pedido, nil
}
