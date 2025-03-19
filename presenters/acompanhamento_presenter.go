package presenters

import (
	"lanchonete/domain/entities"
	"time"
)

type PedidoDTO struct {
Status        entities.StatusPedido `json:"status"`
TempoEstimado time.Duration         `json:"tempoEstimado"`
}

type AcompanhamentoDTO struct {
Pedidos []PedidoDTO `json:"pedidos"`
}

func NewAcompanhamentoDTO(acompanhamento *entities.AcompanhamentoPedido) AcompanhamentoDTO {
	pedidosDTO := make([]PedidoDTO, acompanhamento.Pedidos.Tamanho())
	for i, pedido := range acompanhamento.Pedidos.Listar() {
		tempoEstimado := time.Until(pedido.TimeStamp.Add(acompanhamento.TempoEstimado))
		pedidosDTO[i] = PedidoDTO{
			Status:        pedido.Status,
			TempoEstimado: tempoEstimado,
		}
	}
	return AcompanhamentoDTO{
	Pedidos: pedidosDTO,
}
}