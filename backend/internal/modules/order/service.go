package order

import (
	"context"
	"time"
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
func (s *Service) CreateOrder(
	ctx context.Context,
	orderID string,
	userID string,
) error {

	order := &Order{
		ID:        orderID,
		UserID:    userID,
		Status:    StatusPlaced,
		CreatedAt: time.Now(),
	}

	// type assertion ONLY for now (temporary)
	memRepo, ok := s.repo.(*InMemoryRepository)
	if !ok {
		return nil
	}

	memRepo.Create(order)
	return nil
}
