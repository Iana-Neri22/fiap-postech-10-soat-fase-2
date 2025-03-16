package entities

import (
	"time"
)
type AcompanhamentoPedido struct {
	Pedidos []Pedido
	TempoEstimado time.Duration
	UltimaAtualizacao time.Time
}
