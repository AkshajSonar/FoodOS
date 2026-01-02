package order

import (
	"context"
	"testing"
)

/*
	Mock publisher for tests
	Implements EventPublisher but does nothing
*/
type noopPublisher struct{}

func (n *noopPublisher) Publish(ctx context.Context, event *OrderEvent) error {
	return nil
}

func TestChangeStatus_Success(t *testing.T) {
	repo := NewInMemoryRepository()
	publisher := &noopPublisher{}

	service := NewService(repo, repo, publisher)

	// seed order
	repo.Create(&Order{
		ID:     "order-1",
		UserID: "user-1",
		Status: StatusPlaced,
	})

	err := service.ChangeStatus(
		context.Background(),
		"order-1",
		StatusConfirmed,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestChangeStatus_InvalidTransition(t *testing.T) {
	repo := NewInMemoryRepository()
	publisher := &noopPublisher{}

	service := NewService(repo, repo, publisher)

	repo.Create(&Order{
		ID:     "order-2",
		UserID: "user-2",
		Status: StatusPlaced,
	})

	err := service.ChangeStatus(
		context.Background(),
		"order-2",
		StatusDelivered,
	)

	if err == nil {
		t.Fatal("expected error for invalid transition, got nil")
	}
}
