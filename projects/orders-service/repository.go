package ordersservice

func NewMemoryOrderRepository() *MemoryOrderRepository {
	orders := make(map[string]Order)

	return &MemoryOrderRepository{orders: orders}
}

func (r *MemoryOrderRepository) Save(order Order) error {
	if _, exists := r.orders[order.ID]; exists {
		return ErrDuplicateOrder
	}
	return nil
}

func (r *MemoryOrderRepository) Get(id string) (Order, error) {
	order, ok := r.orders[id]
	if ok != true {
		return Order{}, ErrOrderNotFound
	}
	return order, nil
}
