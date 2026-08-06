package ordersservice

type OrderRequest struct {
	ID    string        `json:"id"`
	Items []ItemRequest `json:"items"`
}

type ItemRequest struct {
	SKU        string `json:"SKU"`
	PriceCents int64  `json:"price_cents"`
	Quantity   int    `json:"quantity"`
}
