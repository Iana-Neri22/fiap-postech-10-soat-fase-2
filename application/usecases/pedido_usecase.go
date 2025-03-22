package usecase

// import (
// 	"context"
// 	"lanchonete/domain/entities"
// 	"lanchonete/domain/repository"
// )

// type pedidoUseCase struct {
// 	pedidoRepo repository.PedidoRepository
// }

// func NewPedidoUseCase(pedidoRepo repository.PedidoRepository) usecases.PedidoUseCase {
// 	return &pedidoUseCase{
// 		pedidoRepo: pedidoRepo,
// 	}
// }

// func (uc *pedidoUseCase) CriarPedido(c context.Context, pedido *entities.Pedido) error {
// 	return uc.pedidoRepo.CriarPedido(c, pedido)
// }

// func (uc *pedidoUseCase) BuscarPedido(c context.Context, identificacao string) (entities.Pedido, error) {
// 	return uc.pedidoRepo.BuscarPedido(c, identificacao)
// }

// func (uc *pedidoUseCase) AtualizarStatusPedido(c context.Context, pedido *entities.Pedido) error {
// 	return uc.pedidoRepo.AtualizarStatusPedido(c, pedido)
// }
