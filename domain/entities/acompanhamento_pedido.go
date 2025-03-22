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

func (a *AcompanhamentoPedido) AdicionarPedido(pedido Pedido) {
	if pedido.Status == Recebido {
		a.Pedidos.Enfileirar(pedido)
		a.UltimaAtualizacao = time.Now()
	}
}

func (a *AcompanhamentoPedido) AtualizarStatusPedido(identificacao string, novoStatus StatusPedido) bool {
	pedidos := a.Pedidos.Listar()
	for i, pedido := range pedidos {
		if pedido.Identificacao == identificacao {
			pedido.Status = novoStatus
			pedido.UltimaAtualizacao = time.Now()

			if novoStatus == Finalizado {
				a.Pedidos = a.removerPedidoNaPosicao(i)
			} else {
				pedidos[i] = pedido
			}

			a.UltimaAtualizacao = time.Now()
			return true
		}
	}
	return false
}

func (a *AcompanhamentoPedido) removerPedidoNaPosicao(posicao int) FilaPedidos {
	pedidos := a.Pedidos.Listar()
	novosPedidos := append(pedidos[:posicao], pedidos[posicao+1:]...)

	novaFila := NewFilaPedidos()
	for _, p := range novosPedidos {
		novaFila.Enfileirar(p)
	}

	return novaFila
}
