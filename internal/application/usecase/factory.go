package usecase

import (
	"lanchonete/internal/infrastructure/repository"
	"lanchonete/usecases"
)

// CasoDeUsoFactory é responsável por criar instâncias de casos de uso
type CasoDeUsoFactory struct {
	repositorioFactory *repository.RepositorioFactory
}

// NovoCasoDeUsoFactory cria uma nova instância de CasoDeUsoFactory
func NovoCasoDeUsoFactory(repositorioFactory *repository.RepositorioFactory) *CasoDeUsoFactory {
	return &CasoDeUsoFactory{
		repositorioFactory: repositorioFactory,
	}
}

// CriarPedidoIncluirUseCase cria uma instância de PedidoIncluirUseCase
func (f *CasoDeUsoFactory) CriarPedidoIncluirUseCase() *usecases.PedidoIncluirUseCase {
	return usecases.NewPedidoIncluirUseCase(f.repositorioFactory.CriarPedidoRepository())
}

// CriarPedidoBuscarPorIdUseCase cria uma instância de PedidoBuscarPorIdUseCase
func (f *CasoDeUsoFactory) CriarPedidoBuscarPorIdUseCase() *usecases.PedidoBuscarPorIdUseCase {
	return usecases.NewPedidoBuscarPorIdUseCase(f.repositorioFactory.CriarPedidoRepository())
}

// CriarPedidoAtualizarStatusUseCase cria uma instância de PedidoAtualizarStatusUseCase
func (f *CasoDeUsoFactory) CriarPedidoAtualizarStatusUseCase() *usecases.PedidoAtualizarStatusUseCase {
	return usecases.NewPedidoAtualizarStatusUseCase(f.repositorioFactory.CriarPedidoRepository())
}

// CriarPedidoListarTodosUseCase cria uma instância de PedidoListarTodosUseCase
func (f *CasoDeUsoFactory) CriarPedidoListarTodosUseCase() *usecases.PedidoListarTodosUseCase {
	return usecases.NewPedidoListarTodosUseCase(f.repositorioFactory.CriarPedidoRepository())
}

// CriarProdutoBuscarPorIdUseCase cria uma instância de ProdutoBuscaPorIdUseCase
func (f *CasoDeUsoFactory) CriarProdutoBuscarPorIdUseCase() *usecases.ProdutoBuscaPorIdUseCase {
	return usecases.NewProdutoBuscaPorIdUseCase(f.repositorioFactory.CriarProdutoRepository())
} 