package ordersservice

import "context"

func NewMemoryOrderRepository() *MemoryOrderRepository {
	orders := make(map[string]Order)

	return &MemoryOrderRepository{orders: orders}
}

func (r *MemoryOrderRepository) Save(ctx context.Context, order Order) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.orders[order.ID]; exists {
		return ErrDuplicateOrder
	}
	r.orders[order.ID] = order
	return nil
}

func (r *MemoryOrderRepository) Get(ctx context.Context, id string) (Order, error) {
	if ctx.Err() != nil {
		return Order{}, ctx.Err()
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	if id == "" {
		return Order{}, ErrInvalidIDOrder
	}
	order, ok := r.orders[id]
	if ok != true {
		return Order{}, ErrOrderNotFound
	}
	return order, nil
}
