package order

import (
	"context"
	"database/sql"
	"time"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

/* -------- ORDERS -------- */

func (r *PostgresRepository) GetByID(
	ctx context.Context,
	id string,
) (*Order, error) {

	row := r.db.QueryRowContext(
		ctx,
		"SELECT id, user_id, status, created_at FROM orders WHERE id=$1",
		id,
	)

	var o Order
	err := row.Scan(&o.ID, &o.UserID, &o.Status, &o.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &o, nil
}

func (r *PostgresRepository) UpdateStatus(
	ctx context.Context,
	id string,
	status OrderStatus,
) error {

	_, err := r.db.ExecContext(
		ctx,
		"UPDATE orders SET status=$1 WHERE id=$2",
		status,
		id,
	)

	return err
}

func (r *PostgresRepository) Create(order *Order) error {
	_, err := r.db.Exec(
		`
		INSERT INTO orders (id, user_id, status, created_at)
		VALUES ($1, $2, $3, $4)
		`,
		order.ID,
		order.UserID,
		order.Status,
		time.Now(),
	)
	return err
}

/* -------- EVENTS -------- */

func (r *PostgresRepository) Save(
	ctx context.Context,
	event *OrderEvent,
) error {

	_, err := r.db.ExecContext(
		ctx,
		`
		INSERT INTO order_events
		(id, order_id, type, old_status, new_status, created_at)
		VALUES ($1,$2,$3,$4,$5,$6)
		`,
		event.ID,
		event.OrderID,
		event.Type,
		event.OldStatus,
		event.NewStatus,
		event.CreatedAt,
	)

	return err
}
