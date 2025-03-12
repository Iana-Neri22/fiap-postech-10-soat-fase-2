package entities

import (
	"time"
)

type StatusPedido string

const (
	Pendente StatusPedido = "Pendente"
	Recebido StatusPedido = "Recebido"
	EmPreparacao StatusPedido = "Em preparação"
	Pronto StatusPedido = "Pronto"
	Finalizado StatusPedido = "Finalizado"
)

type AcompanhamentoPedido struct {
	PedidoID string
	Status StatusPedido
	TempoEstimado time.Duration
	UltimaAtualizacao time.Time
}
