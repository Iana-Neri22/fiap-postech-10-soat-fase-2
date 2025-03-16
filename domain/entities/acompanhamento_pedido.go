package entities

import (
	"time"
)
type AcompanhamentoPedido struct {
	Pedidos FilaPedidos
	TempoEstimado time.Duration
	UltimaAtualizacao time.Time
}
