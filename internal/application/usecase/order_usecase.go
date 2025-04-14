package usecase

import (
	"context"
	"github.com/your-username/your-project/internal/domain/entities"
	"github.com/your-username/your-project/internal/domain/repository"
)

// orderUseCase implements the OrderUseCase interface
type orderUseCase struct {
	orderRepo repository.OrderRepository
}

// NewOrderUseCase creates a new instance of OrderUseCase
func NewOrderUseCase(orderRepo repository.OrderRepository) *orderUseCase {
	return &orderUseCase{
		orderRepo: orderRepo,
	}
}

// CreateOrder implements the CreateOrder method of OrderUseCase
func (uc *orderUseCase) CreateOrder(ctx context.Context, order *entities.Order) error {
	return uc.orderRepo.Create(ctx, order)
}

// GetOrder implements the GetOrder method of OrderUseCase
func (uc *orderUseCase) GetOrder(ctx context.Context, id string) (*entities.Order, error) {
	return uc.orderRepo.GetByID(ctx, id)
}

// UpdateOrder implements the UpdateOrder method of OrderUseCase
func (uc *orderUseCase) UpdateOrder(ctx context.Context, order *entities.Order) error {
	return uc.orderRepo.Update(ctx, order)
}

// DeleteOrder implements the DeleteOrder method of OrderUseCase
func (uc *orderUseCase) DeleteOrder(ctx context.Context, id string) error {
	return uc.orderRepo.Delete(ctx, id)
}

// ListOrders implements the ListOrders method of OrderUseCase
func (uc *orderUseCase) ListOrders(ctx context.Context) ([]*entities.Order, error) {
	return uc.orderRepo.List(ctx)
} 