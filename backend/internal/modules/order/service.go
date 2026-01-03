package order

import (
	"context"
	"time"
)

type Service struct {
	repo       Repository
	eventRepo  EventRepository
	publisher  EventPublisher
}

func NewService(
	r Repository,
	e EventRepository,
	p EventPublisher,
) *Service {
	return &Service{
		repo:      r,
		eventRepo: e,
		publisher: p,
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
	_ = s.publisher.Publish(ctx, &OrderEvent{
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

	if err := s.repo.(interface {
		Create(*Order) error
	}).Create(order); err != nil {
		return err
	}

	_ = s.eventRepo.Save(ctx, &OrderEvent{
		ID:        "evt-" + orderID,
		OrderID:   orderID,
		Type:      EventOrderCreated,
		NewStatus: StatusPlaced,
		CreatedAt: time.Now(),
	})

	_ = s.publisher.Publish(ctx, &OrderEvent{
		ID:        "evt-" + orderID,
		OrderID:   orderID,
		Type:      EventOrderCreated,
		NewStatus: StatusPlaced,
		CreatedAt: time.Now(),
	})

	return nil
}
func (s *Service) GetByID(ctx context.Context, id string) (*Order, error) {
	return s.repo.GetByID(ctx, id)
}
func (s *Service) List(ctx context.Context) ([]*Order, error) {
	return s.repo.List(ctx)
}


