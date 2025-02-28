package interfaces

import (
	"lanchonete/domain/entities"
)

type IClienteGateway interface {
	CriarCliente(*entities.Cliente) error
	BuscarCliente(string) (*entities.Cliente, error)
}