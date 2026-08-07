package ordersservice

import (
	"context"
)

type Handler struct {
	service OrderService
}

func NewHandler(ctx context.Context, service OrderService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(ctx context.Context, service OrderService) error {
	//order := Order{}

	/*
		body := json.NewEncoder(io.Writer(order))

		req := httptest.NewRequestWithContext(
			ctx,
			http.MethodPost,
			"/order",
			body,
		)
	*/
	return ErrDuplicateOrder
}
