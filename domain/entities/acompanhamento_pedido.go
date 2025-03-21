package entities

import (
	"time"
)

const (
	DefaultPreparationTime = 15 * time.Minute
)

type AcompanhamentoPedido struct {
	ID               string
	Pedidos          FilaPedidos
	TempoEstimado    time.Duration
	UltimaAtualizacao time.Time
}

func NewAcompanhamentoPedido() *AcompanhamentoPedido {
	return &AcompanhamentoPedido{
		Pedidos:           NewFilaPedidos(),
		TempoEstimado:     DefaultPreparationTime,
		UltimaAtualizacao: time.Now(),
	}
}

// AdicionarPedido adds a paid order to the queue
func (a *AcompanhamentoPedido) AdicionarPedido(pedido Pedido) {
	// Only add orders with status "Recebido" (paid)
	if pedido.Status == Recebido {
		a.Pedidos.Enfileirar(pedido)
		a.UltimaAtualizacao = time.Now()
	}
}

// AtualizarStatusPedido updates the status of an order in the queue
func (a *AcompanhamentoPedido) AtualizarStatusPedido(identificacao string, novoStatus StatusPedido) bool {
	pedidos := a.Pedidos.Listar()
	for i, pedido := range pedidos {
		if pedido.Identificacao == identificacao {
			// Update the status
			pedido.Status = novoStatus
			pedido.UltimaAtualizacao = time.Now()

			// If status is "Finalizado", remove from queue
			if novoStatus == Finalizado {
				a.Pedidos = a.removerPedidoNaPosicao(i)
			} else {
				// Otherwise update the pedido in the queue
				pedidos[i] = pedido
			}

			a.UltimaAtualizacao = time.Now()
			return true
		}
	}
	return false
}

// removerPedidoNaPosicao removes an order at a specific position
func (a *AcompanhamentoPedido) removerPedidoNaPosicao(posicao int) FilaPedidos {
	pedidos := a.Pedidos.Listar()
	novosPedidos := append(pedidos[:posicao], pedidos[posicao+1:]...)

	novaFila := NewFilaPedidos()
	for _, p := range novosPedidos {
		novaFila.Enfileirar(p)
	}

	return novaFila
}
