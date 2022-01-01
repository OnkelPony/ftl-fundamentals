package creditcard

import "errors"

type card struct {
	number string
}

func New(n string) (*card, error) {
	if n == "" {
		return &card{}, errors.New("invalid card number")
	}
	return &card{
		number: n,
	}, nil
}

func (c *card) SetNumber(n string) error {
	if n == "" {
		return errors.New("invalid card number")
	}
	c.number = n
	return nil
}

func (c card) Number() string {
	return c.number
}
