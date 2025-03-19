package entities

import (
	"time"
)

const (
	DefaultPreparationTime = 15 * time.Minute
)

type AcompanhamentoPedido struct {
	Pedidos           FilaPedidos
	TempoEstimado     time.Duration
	UltimaAtualizacao time.Time
}
