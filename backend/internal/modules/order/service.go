package order

import (
	"context"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) ChangeStatus(
	ctx context.Context,
	orderID string,
	newStatus OrderStatus,
) error {

	order, err := s.repo.GetByID(ctx, orderID)
	if err != nil {
		return err
	}

	if !CanTransition(order.Status, newStatus) {
		return ErrInvalidTransition
	}

	return s.repo.UpdateStatus(ctx, orderID, newStatus)
}
