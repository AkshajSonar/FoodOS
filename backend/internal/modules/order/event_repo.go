package order

import "context"

type EventRepository interface {
	Save(ctx context.Context, event *OrderEvent) error
}
