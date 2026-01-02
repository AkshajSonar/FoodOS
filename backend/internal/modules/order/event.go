package order

import "time"

type OrderEventType string

const (
	EventOrderCreated OrderEventType = "ORDER_CREATED"
	EventStatusChanged OrderEventType = "STATUS_CHANGED"
)

type OrderEvent struct {
	ID        string
	OrderID   string
	Type      OrderEventType
	OldStatus OrderStatus
	NewStatus OrderStatus
	CreatedAt time.Time
}
