package order

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrOrderNotFound = errors.New("order not found")

type InMemoryRepository struct {
	mu     sync.RWMutex
	orders map[string]*Order
	events []*OrderEvent
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		orders: make(map[string]*Order),
		events: []*OrderEvent{},
	}
}

func (r *InMemoryRepository) GetByID(
	ctx context.Context,
	id string,
) (*Order, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	order, ok := r.orders[id]
	if !ok {
		return nil, ErrOrderNotFound
	}

	return order, nil
}

func (r *InMemoryRepository) UpdateStatus(
	ctx context.Context,
	id string,
	status OrderStatus,
) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	order, ok := r.orders[id]
	if !ok {
		return ErrOrderNotFound
	}

	order.Status = status
	return nil
}

//
// helper method (TEMP) to seed orders
//
func (r *InMemoryRepository) Create(order *Order) {
	r.mu.Lock()
	defer r.mu.Unlock()

	order.CreatedAt = time.Now()
	r.orders[order.ID] = order
}

func (r *InMemoryRepository) Save(
	ctx context.Context,
	event *OrderEvent,
) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.events = append(r.events, event)
	return nil
}

