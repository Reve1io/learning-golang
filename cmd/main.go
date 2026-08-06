package main

import (
	"context"
	"fmt"
	ordersservice "test/projects/orders-service"
	"time"
)

func main() {
	orders := []ordersservice.Order{
		{
			ID: "1",
			Items: []ordersservice.Item{
				{SKU: "milk", PriceCents: 199, Quantity: 2},
				{SKU: "bread", PriceCents: 500, Quantity: 1},
			},
		},
		{
			ID: "1",
			Items: []ordersservice.Item{
				{SKU: "Strawberry", PriceCents: 199, Quantity: 2},
				{SKU: "Apple", PriceCents: 500, Quantity: 1},
			},
		},
	}
	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancel()

	repo := ordersservice.NewMemoryOrderRepository()
	service := ordersservice.NewOrderService(repo)

	for _, order := range orders {
		push := service.Create(ctx, order)
		fmt.Println(push)

		getedOrder, err := service.GetOrderByID(ctx, order.ID)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Order: %v\n", getedOrder)

		totalPrice := service.CountOrder(order)
		fmt.Printf("Total order: %v\n", totalPrice)
	}
}
