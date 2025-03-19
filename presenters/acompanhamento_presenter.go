package presenters

import (
	"lanchonete/domain/entities"
	"time"
)

type PedidoDTO struct {
	Identificacao string                `json:"identificacao"`
	Status        entities.StatusPedido `json:"status"`
	TempoEstimado time.Duration         `json:"tempoEstimado"`
}

type AcompanhamentoDTO struct {
	Pedidos []PedidoDTO `json:"pedidos"`
}

func NewAcompanhamentoDTO(acompanhamento *entities.AcompanhamentoPedido) AcompanhamentoDTO {
	pedidosDTO := make([]PedidoDTO, acompanhamento.Pedidos.Tamanho())
	for i, pedido := range acompanhamento.Pedidos.Listar() {
		var tempoEstimado time.Duration

		if pedido.Status == entities.EmPreparacao {
			startTime := pedido.UltimaAtualizacao
			elapsedTime := time.Since(startTime)
			tempoEstimado := entities.DefaultPreparationTime - elapsedTime

			if tempoEstimado < 0 {
				tempoEstimado = 0
			}
		} else {
			tempoEstimado = time.Duration(0)
		}
		pedidosDTO[i] = PedidoDTO{
			Status:        pedido.Status,
			TempoEstimado: tempoEstimado,
		}
	}
	return AcompanhamentoDTO{
	Pedidos: pedidosDTO,
}
}