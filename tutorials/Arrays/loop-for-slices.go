package arrays

func FilterExpensiveOrders(orders []int, limit int) []int {

	if len(orders) == 0 {
		return orders
	}

	ordersHighSumm := make([]int, 0)

	for _, order := range orders {

		if order > limit {
			ordersHighSumm = append(ordersHighSumm, order)
		} else if order < limit {
			continue
		}
	}

	return ordersHighSumm
}
