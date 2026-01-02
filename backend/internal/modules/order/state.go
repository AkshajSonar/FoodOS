package order

import "errors"

var ErrInvalidTransition = errors.New("invalid order state transition")

var allowedTransitions = map[OrderStatus][]OrderStatus{
	StatusPlaced: {
		StatusConfirmed,
		StatusCancelled,
	},
	StatusConfirmed: {
		StatusPreparing,
		StatusCancelled,
	},
	StatusPreparing: {
		StatusPicked,
	},
	StatusPicked: {
		StatusDelivered,
	},
}

func CanTransition(from, to OrderStatus) bool {
	nextStates, ok := allowedTransitions[from]
	if !ok {
		return false
	}

	for _, s := range nextStates {
		if s == to {
			return true
		}
	}
	return false
}
