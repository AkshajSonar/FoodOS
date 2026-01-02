package order

import (
	"context"
	"time"
)

type Service struct {
	repo      Repository
	eventRepo EventRepository
}

func NewService(r Repository, e EventRepository) *Service {
	return &Service{
		repo:      r,
		eventRepo: e,
	}
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

	oldStatus := order.Status

	if err := s.repo.UpdateStatus(ctx, orderID, newStatus); err != nil {
		return err
	}

	// record event
	_ = s.eventRepo.Save(ctx, &OrderEvent{
		ID:        "evt-" + orderID + "-" + string(newStatus),
		OrderID:   orderID,
		Type:      EventStatusChanged,
		OldStatus: oldStatus,
		NewStatus: newStatus,
		CreatedAt: time.Now(),
	})

	return nil
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

	memRepo, ok := s.repo.(*InMemoryRepository)
	if !ok {
		return nil
	}

	memRepo.Create(order)

	// record event
	_ = s.eventRepo.Save(ctx, &OrderEvent{
		ID:        "evt-" + orderID,
		OrderID:   orderID,
		Type:      EventOrderCreated,
		NewStatus: StatusPlaced,
		CreatedAt: time.Now(),
	})

	return nil
}

