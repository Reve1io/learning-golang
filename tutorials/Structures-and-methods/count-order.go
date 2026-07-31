package structuresandmethods

import (
	"errors"
	"fmt"
)

type Item struct {
	SKU        string
	PriceCents int64
	Quantity   int
}

type Order struct {
	Items []Item
}

var (
	ErrEmptyOrder  = errors.New("order is empty")
	ErrInvalidItem = errors.New("invalid item")
)

func ProcessOrder(order Order) (int64, error) {
	var total int64 = 0
	var err error

	err = order.Validation()
	if err != nil {
		return total, err
	} else {
		total = order.TotalCents()
	}

	return total, err
}

func (o Order) TotalCents() int64 {
	var total int64

	for _, item := range o.Items {
		sum := item.PriceCents * int64(item.Quantity)
		total += sum
	}

	return total
}

func (o Order) IsEmpty() bool {
	if len(o.Items) == 0 {
		return true
	}

	return false
}

func (o Order) Validation() error {

	var err error

	if len(o.Items) == 0 {
		return ErrEmptyOrder
	}

	for index, item := range o.Items {
		switch {
		case item.SKU == "":
			return fmt.Errorf("item %d: %w", index, ErrInvalidItem)
		case item.PriceCents < 0:
			return fmt.Errorf("item %d: %w", index, ErrInvalidItem)
		case item.Quantity <= 0:
			return fmt.Errorf("item %d: %w", index, ErrInvalidItem)
		}
	}

	return err
}
