package order

import "context"

type Repository interface {
	GetByID(ctx context.Context, id string) (*Order, error)
	UpdateStatus(ctx context.Context, id string, status OrderStatus) error
	List(ctx context.Context) ([]*Order, error)
}
