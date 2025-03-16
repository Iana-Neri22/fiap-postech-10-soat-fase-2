package service

import (
	"sync"
	"lanchonete/domain/entities"
)

type FilaPedidos struct {
	pedidos []entities.Pedido
	mu      sync.Mutex
}

func NewFila() *FilaPedidos {
	return &FilaPedidos{}
}

func (f *FilaPedidos) Enfileirar(p entities.Pedido) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.pedidos = append(f.pedidos, p)
}

func (f *FilaPedidos) Desenfileirar() (entities.Pedido, bool) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if len(f.pedidos) == 0 {
		return entities.Pedido{}, false
	}

	pedido := f.pedidos[0]
	f.pedidos = f.pedidos[1:]
	return pedido, true
}

func (f *FilaPedidos) Tamanho() int {
	f.mu.Lock()
	defer f.mu.Unlock()

	return len(f.pedidos)
}
