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
	ID        string      `json:"id"`
	UserID    string      `json:"userId"`
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time   `json:"createdAt"`
}

