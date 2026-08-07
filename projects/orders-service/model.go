package ordersservice

import "sync"

type Order struct {
	ID    string `json:"id"`
	Items []Item `json:"items"`
}

type Item struct {
	SKU        string `json:"sku"`
	PriceCents int64  `json:"price_cents"`
	Quantity   int    `json:"quantity"`
}

type MemoryOrderRepository struct {
	mu     sync.RWMutex
	orders map[string]Order
}
