package repository

import (
	"lanchonete/gateways"
	"lanchonete/internal/domain/repository"
	"lanchonete/infra/database/mongo"
)

// RepositorioFactory é responsável por criar instâncias de repositórios
type RepositorioFactory struct {
	db mongo.Database
}

// NovoRepositorioFactory cria uma nova instância de RepositorioFactory
func NovoRepositorioFactory(db mongo.Database) *RepositorioFactory {
	return &RepositorioFactory{
		db: db,
	}
}

// CriarPedidoRepository cria uma instância de PedidoRepository
func (f *RepositorioFactory) CriarPedidoRepository() repository.PedidoRepository {
	return gateways.NewPedidoGateway(f.db, mongo.CollectionPedido)
}

// CriarProdutoRepository cria uma instância de ProdutoRepository
func (f *RepositorioFactory) CriarProdutoRepository() repository.ProdutoRepository {
	return gateways.NewProdutoGateway(f.db, mongo.CollectionProduto)
}

// CriarClienteRepository cria uma instância de ClienteRepository
func (f *RepositorioFactory) CriarClienteRepository() repository.ClienteRepository {
	return gateways.NewClienteRepository(f.db, mongo.CollectionCliente)
}

// CriarPagamentoRepository cria uma instância de PagamentoRepository
func (f *RepositorioFactory) CriarPagamentoRepository() repository.PagamentoRepository {
	return gateways.NewPagamentoRepository(f.db, mongo.CollectionPagamento)
}

// CriarAcompanhamentoRepository cria uma instância de AcompanhamentoRepository
func (f *RepositorioFactory) CriarAcompanhamentoRepository() repository.AcompanhamentoRepository {
	return gateways.NewAcompanhamentoRepository(f.db, mongo.CollectionAcompanhamento)
} 