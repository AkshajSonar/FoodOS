package order

import "testing"

func TestValidOrderTransitions(t *testing.T) {
	tests := []struct {
		from OrderStatus
		to   OrderStatus
	}{
		{StatusPlaced, StatusConfirmed},
		{StatusPlaced, StatusCancelled},
		{StatusConfirmed, StatusPreparing},
		{StatusPreparing, StatusPicked},
		{StatusPicked, StatusDelivered},
	}

	for _, tt := range tests {
		if !CanTransition(tt.from, tt.to) {
			t.Errorf("expected valid transition from %s to %s", tt.from, tt.to)
		}
	}
}

func TestInvalidOrderTransitions(t *testing.T) {
	tests := []struct {
		from OrderStatus
		to   OrderStatus
	}{
		{StatusPlaced, StatusDelivered},
		{StatusConfirmed, StatusPicked},
		{StatusDelivered, StatusPreparing},
	}

	for _, tt := range tests {
		if CanTransition(tt.from, tt.to) {
			t.Errorf("expected invalid transition from %s to %s", tt.from, tt.to)
		}
	}
}
