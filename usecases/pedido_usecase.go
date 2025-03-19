package usecase

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
<<<<<<< HEAD
=======
	"lanchonete/domain/usecases"
>>>>>>> iana
)

type pedidoUseCase struct {
	pedidoRepository repository.PedidoRepository
}

<<<<<<< HEAD
func NewPedidoUseCase(pedidoRepository repository.PedidoRepository) repository.PedidoRepository {
=======
func NewPedidoUseCase(pedidoRepository repository.PedidoRepository) usecases.PedidoUseCase {
>>>>>>> iana
	return &pedidoUseCase{
		pedidoRepository: pedidoRepository,
	}
}

func (puc *pedidoUseCase) CriarPedido(c context.Context, pedido *entities.Pedido) error {
	return puc.pedidoRepository.CriarPedido(c, pedido)
}

<<<<<<< HEAD
func (puc *pedidoUseCase) BuscarPedido(c context.Context, identificacao string) (*entities.Pedido, error) {
=======
func (puc *pedidoUseCase) BuscarPedido(c context.Context, identificacao string) (entities.Pedido, error) {
>>>>>>> iana
	return puc.pedidoRepository.BuscarPedido(c, identificacao)
}

func (puc *pedidoUseCase) AtualizarStatusPedido(c context.Context, identificacao string, status string) error {
	return puc.pedidoRepository.AtualizarStatusPedido(c, identificacao, status)
}
