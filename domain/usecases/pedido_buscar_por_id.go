package usecases

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
)

type PedidoBuscarPorIdUseCase struct {
	pedidoRepository repository.PedidoRepository
}

func NewPedidoBuscarPorIdUseCase(pedidoRepository repository.PedidoRepository) *PedidoBuscarPorIdUseCase {
	return &PedidoBuscarPorIdUseCase{
		pedidoRepository: pedidoRepository,
	}
}

// <<<<<<< HEAD
func (pduc *PedidoBuscarPorIdUseCase) Run(c context.Context, identificacao string) (entities.Pedido, error) {
	pedido, err := pduc.pedidoRepository.BuscarPedido(c, identificacao)
	if err != nil {
		return entities.Pedido{}, err
// =======
// func (pduc *PedidoBuscarPorIdUseCase) Run(c context.Context, identificacao string) (*entities.Pedido, error) {

// 	pedido, err := pduc.pedidoRepository.BuscarPedido(c, identificacao)
// 	if err != nil {
// 		return nil, err
// >>>>>>> Bruno
	}
	return pedido, nil
}
