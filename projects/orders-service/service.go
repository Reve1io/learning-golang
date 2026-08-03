package ordersservice

import (
	"errors"
	"fmt"
)

type OrderRepository interface {
	Save(order Order) error
	Get(id string) (Order, error)
}

type OrderService struct {
	repo OrderRepository
}

var (
	ErrEmptyOrder     = errors.New("order is empty")
	ErrInvalidItem    = errors.New("invalid item")
	ErrOrderNotFound  = errors.New("order not found")
	ErrDuplicateOrder = errors.New("duplicate order")
)

func NewOrderService(repo OrderRepository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) Create(order Order) error {
	if order.IsEmpty() {
		return ErrEmptyOrder
	}

	if err := order.Validation(); err != nil {
		return fmt.Errorf("error: %w", err)
	}

	_, err := s.repo.Get(order.ID)
	if err == nil {
		return ErrDuplicateOrder
	}

	return s.repo.Save(order)
}

func (s *OrderService) GetOrderByID(id string) (Order, error) {
	o, err := s.repo.Get(id)
	if err != nil {
		return Order{}, err
	}
	return o, nil
}

func (s *OrderService) CountOrder(order Order) int64 {
	totalOrder := order.TotalCents()

	return totalOrder
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
