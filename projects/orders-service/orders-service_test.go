package ordersservice_test

import (
	"errors"
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
		name   string
		input1 []ordersservice.Order
		input2 []ordersservice.Order
		want   error
	}{
		{
			name: "Two item case",
			input1: []ordersservice.Order{
				{
					ID: "1",
					Items: []ordersservice.Item{
						{SKU: "milk", PriceCents: 199, Quantity: 2},
						{SKU: "bread", PriceCents: 500, Quantity: 1},
					},
				},
			},
			input2: []ordersservice.Order{
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
			input1: []ordersservice.Order{
				{
					ID: "3",
					Items: []ordersservice.Item{
						{SKU: "milk", PriceCents: 199, Quantity: 2},
						{SKU: "bread", PriceCents: 500, Quantity: 1},
					},
				},
			},
			input2: []ordersservice.Order{
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
			repo := ordersservice.NewMemoryOrderRepository()
			service := ordersservice.NewOrderService(repo)

			for _, inp := range tt.input {
				got := service.Create(inp)
				if !errors.Is(got, tt.want) {
					t.Errorf("Saving products(%v) = %v; want %v", tt.input, got, tt.want)
				}
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
	}

	repo := ordersservice.NewMemoryOrderRepository()
	service := ordersservice.NewOrderService(repo)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, inp := range tt.input {
				_, got := service.GetOrderByID(inp.ID)
				if !errors.Is(got, tt.want) {
					t.Errorf("Get products(%v) = %v; want %v", tt.input, got, tt.want)
				}
			}
		})
	}
}
