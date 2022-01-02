package order

import (
	"errors"
)

type Order struct {
	id string
}

func New(id string) (*Order, error) {
	if id == "" {
		return nil, errors.New("invalid order id")
	}
	return &Order{id: id}, nil
}

func (o Order) ID() string {
	return o.id
}

func (o *Order) SetID(id string) error {
	if id == "" {
		return errors.New("invalid order id")
	}
	o.id = id
	return nil
}
