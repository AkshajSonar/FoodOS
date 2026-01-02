package order

import (
	"context"
	"testing"
)

func TestChangeStatus_Success(t *testing.T) {
	repo := NewInMemoryRepository()
	service := NewService(repo, repo)

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
	service := NewService(repo, repo)

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
