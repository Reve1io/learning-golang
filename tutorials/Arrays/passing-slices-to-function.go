package arrays

import "slices"

func AddSiscount(prices []int, discount int) []int {

	if discount > 100 || discount < 0 {
		return prices
	}

	newPrices := slices.Clone(prices)

	for i, price := range newPrices {

		price = price - (price * discount / 100)
		newPrices[i] = price
	}

	return newPrices
}
