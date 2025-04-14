package usecase

import (
	"context"
	"github.com/your-username/your-project/internal/domain/entities"
)

// OrderUseCase defines the interface for order business operations
type OrderUseCase interface {
	CreateOrder(ctx context.Context, order *entities.Order) error
	GetOrder(ctx context.Context, id string) (*entities.Order, error)
	UpdateOrder(ctx context.Context, order *entities.Order) error
	DeleteOrder(ctx context.Context, id string) error
	ListOrders(ctx context.Context) ([]*entities.Order, error)
} 