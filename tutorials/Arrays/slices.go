package arrays

func AddDiscount(prices []float64, discount float64) []float64 {

	if len(prices) == 0 {
		return prices
	}

	latestNum := prices[len(prices)-1]

	percent := latestNum / 100 * discount

	prices[len(prices)-1] = latestNum - percent

	return prices
}
