package order

import "time"

type OrderStatus string

const (
	StatusPlaced    OrderStatus = "PLACED"
	StatusConfirmed OrderStatus = "CONFIRMED"
	StatusPreparing OrderStatus = "PREPARING"
	StatusPicked    OrderStatus = "PICKED"
	StatusDelivered OrderStatus = "DELIVERED"
	StatusCancelled OrderStatus = "CANCELLED"
)

type Order struct {
	ID        string
	UserID    string
	Status    OrderStatus
	CreatedAt time.Time
}
