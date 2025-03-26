package presenters

import (
	"lanchonete/domain/entities"
	"time"
)

type PedidoDTO struct {
	ClienteCPF     string                `json:"clienteCPF"`
	Personalizacao string                `json:"personalizacao"`
	ID             string                `json:"id"`
	Identificacao  string                `json:"identificacao"`
	Produtos       []string              `json:"produtos"`
	Status         entities.StatusPedido `json:"status"`
	TempoEstimado  time.Duration         `json:"tempoEstimado"`
}
