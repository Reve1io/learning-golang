package structuresandmethods_test

import (
	"errors"
	structuresandmethods "test/tutorials/Structures-and-methods"
	"testing"
)

func TestCountOrder(t *testing.T) {
	tests := []struct {
		name  string
		input structuresandmethods.Order
		want  int64
	}{
		{
			name: "Normal default case (two positions)",
			input: structuresandmethods.Order{
				Items: []structuresandmethods.Item{
					{SKU: "milk", PriceCents: 199, Quantity: 2},
					{SKU: "bread", PriceCents: 500, Quantity: 1},
				},
			},
			want: 898,
		},
		{
			name: "Normal default case (one position)",
			input: structuresandmethods.Order{
				Items: []structuresandmethods.Item{
					{SKU: "milk", PriceCents: 299, Quantity: 2},
				},
			},
			want: 598,
		},
		{
			name: "Quantity value 0 case",
			input: structuresandmethods.Order{
				Items: []structuresandmethods.Item{
					{SKU: "milk", PriceCents: 199, Quantity: 0},
				},
			},
			want: 0,
		},
		{
			name: "Struct is empty case",
			input: structuresandmethods.Order{
				Items: []structuresandmethods.Item{
					{},
				},
			},
			want: 0,
		},
		{
			name: "Struct is empty case",
			input: structuresandmethods.Order{
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
		input structuresandmethods.Order
		want  error
	}{
		{
			name: "Default normal item case",
			input: structuresandmethods.Order{
				Items: []structuresandmethods.Item{
					{SKU: "milk", PriceCents: 199, Quantity: 2},
					{SKU: "bread", PriceCents: 500, Quantity: 1},
				},
			},
			want: nil,
		},
		{
			name: "Order without items",
			input: structuresandmethods.Order{
				Items: []structuresandmethods.Item{},
			},
			want: structuresandmethods.ErrEmptyOrder,
		},
		{
			name: "Order without SKU",
			input: structuresandmethods.Order{
				Items: []structuresandmethods.Item{
					{PriceCents: 199, Quantity: 0},
				},
			},
			want: structuresandmethods.ErrInvalidItem,
		},
		{
			name: "Price is negative number",
			input: structuresandmethods.Order{
				Items: []structuresandmethods.Item{
					{SKU: "milk", PriceCents: -199, Quantity: 2},
				},
			},
			want: structuresandmethods.ErrInvalidItem,
		},
		{
			name: "Quantity is null",
			input: structuresandmethods.Order{
				Items: []structuresandmethods.Item{
					{SKU: "milk", PriceCents: 199, Quantity: 0},
				},
			},
			want: structuresandmethods.ErrInvalidItem,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Validation()
			if !errors.Is(got, tt.want) {
				t.Errorf("CountProducts(%v) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}
