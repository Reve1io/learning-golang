package ordersservice

import "sync"

type Item struct {
	SKU        string
	PriceCents int64
	Quantity   int
}

type Order struct {
	ID    string
	Items []Item
}

type MemoryOrderRepository struct {
	mu     sync.RWMutex
	orders map[string]Order
}
