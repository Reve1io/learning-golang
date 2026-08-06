package ordersservice_test

import (
	"context"
	"errors"
	"sync"
	ordersservice "test/projects/orders-service"
	"testing"
)

func TestCountOrder(t *testing.T) {
	tests := []struct {
		name  string
		input ordersservice.Order
		want  int64
	}{
		{
			name: "Normal default case (two positions)",
			input: ordersservice.Order{
				Items: []ordersservice.Item{
					{SKU: "milk", PriceCents: 199, Quantity: 2},
					{SKU: "bread", PriceCents: 500, Quantity: 1},
				},
			},
			want: 898,
		},
		{
			name: "Normal default case (one position)",
			input: ordersservice.Order{
				Items: []ordersservice.Item{
					{SKU: "milk", PriceCents: 299, Quantity: 2},
				},
			},
			want: 598,
		},
		{
			name: "Quantity value 0 case",
			input: ordersservice.Order{
				Items: []ordersservice.Item{
					{SKU: "milk", PriceCents: 199, Quantity: 0},
				},
			},
			want: 0,
		},
		{
			name: "Struct is empty case",
			input: ordersservice.Order{
				Items: []ordersservice.Item{
					{},
				},
			},
			want: 0,
		},
		{
			name: "Struct is empty case",
			input: ordersservice.Order{
				Items: nil,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCount := tt.input.TotalCents()
			if gotCount != tt.want {
				t.Errorf("CountProducts(%v) = %v; want %v", tt.input, gotCount, tt.want)
			}
		})
	}
}

func TestValidateOrder(t *testing.T) {
	tests := []struct {
		name  string
		input ordersservice.Order
		want  error
	}{
		{
			name: "Default normal item case",
			input: ordersservice.Order{
				Items: []ordersservice.Item{
					{SKU: "milk", PriceCents: 199, Quantity: 2},
					{SKU: "bread", PriceCents: 500, Quantity: 1},
				},
			},
			want: nil,
		},
		{
			name: "Order without items",
			input: ordersservice.Order{
				Items: []ordersservice.Item{},
			},
			want: ordersservice.ErrEmptyOrder,
		},
		{
			name: "Order without SKU",
			input: ordersservice.Order{
				Items: []ordersservice.Item{
					{PriceCents: 199, Quantity: 0},
				},
			},
			want: ordersservice.ErrInvalidItem,
		},
		{
			name: "Price is negative number",
			input: ordersservice.Order{
				Items: []ordersservice.Item{
					{SKU: "milk", PriceCents: -199, Quantity: 2},
				},
			},
			want: ordersservice.ErrInvalidItem,
		},
		{
			name: "Quantity is null",
			input: ordersservice.Order{
				Items: []ordersservice.Item{
					{SKU: "milk", PriceCents: 199, Quantity: 0},
				},
			},
			want: ordersservice.ErrInvalidItem,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Validation()
			if !errors.Is(got, tt.want) {
				t.Errorf("Validate products(%v) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestSavingOrder(t *testing.T) {
	tests := []struct {
		name  string
		input []ordersservice.Order
		want  error
	}{
		{
			name: "Two item case",
			input: []ordersservice.Order{
				{
					ID: "1",
					Items: []ordersservice.Item{
						{SKU: "milk", PriceCents: 199, Quantity: 2},
						{SKU: "bread", PriceCents: 500, Quantity: 1},
					},
				},
				{
					ID: "2",
					Items: []ordersservice.Item{
						{SKU: "milk", PriceCents: 199, Quantity: 2},
						{SKU: "bread", PriceCents: 500, Quantity: 1},
					},
				},
			},
			want: nil,
		},
		{
			name: "Duplicate item case",
			input: []ordersservice.Order{
				{
					ID: "3",
					Items: []ordersservice.Item{
						{SKU: "milk", PriceCents: 199, Quantity: 2},
						{SKU: "bread", PriceCents: 500, Quantity: 1},
					},
				},
				{
					ID: "3",
					Items: []ordersservice.Item{
						{SKU: "milk", PriceCents: 199, Quantity: 2},
						{SKU: "bread", PriceCents: 500, Quantity: 1},
					},
				},
			},
			want: ordersservice.ErrDuplicateOrder,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			repo := ordersservice.NewMemoryOrderRepository()
			service := ordersservice.NewOrderService(repo)

			firstOrder := tt.input[0]
			nextOrder := tt.input[1]

			if err := service.Create(ctx, firstOrder); err != nil {
				t.Fatalf("first Create() error = %v, want = %v", err, tt.want)
			}

			err := service.Create(ctx, nextOrder)

			if !errors.Is(err, tt.want) {
				t.Fatalf("Next Create() error = %v, want = %v", err, tt.want)
			}
		})
	}
}

func TestGetOrder(t *testing.T) {
	tests := []struct {
		name  string
		input []ordersservice.Order
		ID    string
		want  error
	}{
		{
			name: "Unknown ID case",
			input: []ordersservice.Order{
				{
					ID: "6",
					Items: []ordersservice.Item{
						{SKU: "milk", PriceCents: 199, Quantity: 2},
						{SKU: "bread", PriceCents: 500, Quantity: 1},
					},
				},
			},
			ID:   "13",
			want: ordersservice.ErrOrderNotFound,
		},
		{
			name: "Default ID case",
			input: []ordersservice.Order{
				{
					ID: "6",
					Items: []ordersservice.Item{
						{SKU: "milk", PriceCents: 199, Quantity: 2},
						{SKU: "bread", PriceCents: 500, Quantity: 1},
					},
				},
			},
			ID:   "6",
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			repo := ordersservice.NewMemoryOrderRepository()
			service := ordersservice.NewOrderService(repo)

			for _, inp := range tt.input {
				if err := service.Create(ctx, inp); err != nil {
					t.Fatalf("Create() error = %v", err)
				}
				gotOrder, err := service.GetOrderByID(ctx, tt.ID)

				if !errors.Is(err, tt.want) {
					t.Fatalf("GetOrderByID() error = %v; want %v", err, tt.want)
				}

				if tt.want == nil && gotOrder.ID != tt.ID {
					t.Errorf("GetOrderByID() ID = %q; want %q", gotOrder.ID, tt.ID)
				}
			}
		})
	}
}

func TestContextCanceled(t *testing.T) {
	order := ordersservice.Order{
		ID: "6",
		Items: []ordersservice.Item{
			{SKU: "milk", PriceCents: 199, Quantity: 2},
			{SKU: "bread", PriceCents: 500, Quantity: 1},
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	repo := ordersservice.NewMemoryOrderRepository()
	service := ordersservice.NewOrderService(repo)

	err := service.Create(ctx, order)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("got %v; want context.Canceled", err)
	}

	foundOrder, err := service.GetOrderByID(ctx, order.ID)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("got %v; want context.Canceled", err)
	}

	t.Logf("Order: %v", foundOrder)
}

func TestConcurrency(t *testing.T) {
	order := ordersservice.Order{
		ID: "6",
		Items: []ordersservice.Item{
			{SKU: "milk", PriceCents: 199, Quantity: 2},
			{SKU: "bread", PriceCents: 500, Quantity: 1},
		},
	}

	type Result struct {
		Success bool
		ErrText error
	}

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo := ordersservice.NewMemoryOrderRepository()

	resultChan := make(chan Result, 100)

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			_, _ = repo.Get(ctx, "1")
			err := repo.Save(ctx, order)

			var success bool

			if err != nil {
				success = false
			} else {
				success = true
			}

			resultChan <- Result{Success: success, ErrText: err}
		}()
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	successCount, dublicateCount := 0, 0

	for res := range resultChan {
		switch {
		case res.ErrText == nil:
			successCount += 1
		case errors.Is(res.ErrText, ordersservice.ErrDuplicateOrder):
			dublicateCount++
		default:
			t.Errorf("unexpected error: %v", res.ErrText)
		}
	}

	if successCount != 1 {
		t.Errorf("Success save = %v; want = 1", successCount)
	}

	if dublicateCount != 99 {
		t.Errorf("dublicate error = %v; want = 99", dublicateCount)
	}

}
