package entities

import (
	"errors"
	"strings"
	"time"
	"sync"
)

type StatusPedido string

const (
	Pendente StatusPedido = "Pendente"
	Recebido StatusPedido = "Recebido"
	EmPreparacao StatusPedido = "Em preparação"
	Pronto StatusPedido = "Pronto"
	Finalizado StatusPedido = "Finalizado"
)

type Pedido struct {
	ID  		string
	Categoria 	string
	Status   	StatusPedido
	TimeStamp	time.Time
	UltimaAtualizacao time.Time
	mu *sync.Mutex
}

func PedidoNew(ID string, Categoria string, Status StatusPedido) (*Pedido, error) {

	if  strings.TrimSpace(ID) == "" || strings.TrimSpace(Categoria) == "" || strings.TrimSpace(string(Status)) == "" {
		return nil, errors.New("nenhum dos campos podem estar em branco")
	}

	return &Pedido{
		ID:  ID,
		Categoria: Categoria,
		Status:   Status,
		TimeStamp: time.Now(),
		mu: &sync.Mutex{},
	}, nil
}

func (a *Pedido) UpdateStatus(status StatusPedido) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Status = status
	a.UltimaAtualizacao = time.Now()
}