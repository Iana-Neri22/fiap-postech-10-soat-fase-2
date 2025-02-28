package service

import (
	"context"
	"lanchonete/domain/entities"
)

const (
	CollectionCliente = "cliente"
)


type ClienteUseCase interface {
	CriarCliente(C context.Context, cliente *entities.Cliente) error
	BuscarCliente(C context.Context, CPF string) (entities.Cliente, error)
}

