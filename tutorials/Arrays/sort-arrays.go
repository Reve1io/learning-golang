package arrays

import (
	"slices"
	"test/model"
)

func SortOrderByCustomerID(orders []model.Order) []model.Order {

	sortedOrders := slices.Clone(orders)
	slices.SortFunc(sortedOrders, func(order1, order2 model.Order) int {
		if order1.CustomerID != order2.CustomerID {
			return order1.CustomerID - order2.CustomerID
		}

		return order1.Price - order2.Price
	})

	return sortedOrders
}
