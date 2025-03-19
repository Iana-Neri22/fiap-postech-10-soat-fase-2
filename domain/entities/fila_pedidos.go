package entities

import (
	"sort"
	"sync"
)

type FilaPedidos struct {
	pedidos []Pedido
	mu      *sync.Mutex
}

func NewFila() *FilaPedidos {
	return &FilaPedidos{
		pedidos: make([]Pedido, 0),
		mu:      &sync.Mutex{},
	}
}

func (f *FilaPedidos) Enfileirar(p Pedido) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.pedidos = append(f.pedidos, p)

	sort.Slice(f.pedidos, func(i, j int) bool {
		return f.pedidos[i].TimeStamp.Before(f.pedidos[j].TimeStamp)
	})
}

func (f *FilaPedidos) Desenfileirar() (Pedido, bool) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if len(f.pedidos) == 0 {
		return Pedido{}, false
	}

	pedido := f.pedidos[0]
	f.pedidos = f.pedidos[1:]
	return pedido, true
}

func (f *FilaPedidos) RemoverPedido(identificacao string) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	for i, pedido := range f.pedidos {
		if pedido.Identificacao == identificacao {
			// Remove the pedido at position i
			f.pedidos = append(f.pedidos[:i], f.pedidos[i+1:]...)
			return true
		}
	}
	return false
}

func (f *FilaPedidos) Listar() []Pedido {
	f.mu.Lock()
	defer f.mu.Unlock()

	result := make([]Pedido, len(f.pedidos))
	copy(result, f.pedidos)
	return result
}

func (f *FilaPedidos) Tamanho() int {
	f.mu.Lock()
	defer f.mu.Unlock()

	return len(f.pedidos)
}
