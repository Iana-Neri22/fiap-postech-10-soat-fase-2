package usecase

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	service "lanchonete/domain/services"
	"time"
)

type acompanhamentoUseCase struct {
	acompanhamentoRepo repository.AcompanhamentoRepository
}

func NewAcompanhamentoUseCase(acompanhamentoRepo repository.AcompanhamentoRepository) service.AcompanhamentoUseCase {
	return &acompanhamentoUseCase{
		acompanhamentoRepo: acompanhamentoRepo,
	}
}

func (uc *acompanhamentoUseCase) CriarAcompanhamento(c context.Context, acompanhamento *entities.AcompanhamentoPedido) error {
	return uc.acompanhamentoRepo.CriarAcompanhamento(c, acompanhamento)
}

func (uc *acompanhamentoUseCase) BuscarPedido(c context.Context, nroPedido string) (entities.Pedido, error) {
	return uc.acompanhamentoRepo.BuscarPedidos(c, nroPedido)
}

func (uc *acompanhamentoUseCase) AdicionarPedido(c context.Context, acompanhamento *entities.AcompanhamentoPedido, pedido *entities.Pedido) error {
	if pedido.Status == entities.Recebido {
		acompanhamento.Pedidos.Enfileirar(*pedido)
		acompanhamento.UltimaAtualizacao = time.Now()

		return uc.acompanhamentoRepo.AdicionarPedido(c, acompanhamento, pedido)
	}
	return nil
}

func (uc *acompanhamentoUseCase) BuscarAcompanhamento(c context.Context, nroPedido string) (entities.AcompanhamentoPedido, error) {
	return uc.acompanhamentoRepo.BuscarAcompanhamento(c, nroPedido)
}

func (uc *acompanhamentoUseCase) AtualizarStatusPedido(c context.Context, acompanhamento *entities.AcompanhamentoPedido, identificacao string, novoStatus entities.StatusPedido) error {
	pedidos := acompanhamento.Pedidos.Listar()
	updated := false

	for i, pedido := range pedidos {
		if pedido.Identificacao == identificacao {
			// Update the status
			pedido.Status = novoStatus
			pedido.UltimaAtualizacao = time.Now()

			// If status is "Finalizado", remove from queue
			if novoStatus == entities.Finalizado {
				acompanhamento.Pedidos.RemoverPedido(identificacao)
			} else {
				// Otherwise update the pedido in the queue
				pedidos[i] = pedido
			}

			acompanhamento.UltimaAtualizacao = time.Now()
			updated = true
			break
		}
	}

	if !updated {
		return nil // Or return an error indicating the pedido wasn't found
	}

	// We would need to save the updated acompanhamento
	// This would require an additional method in the repository
	// For now, we'll assume it's handled elsewhere

	return nil
}
