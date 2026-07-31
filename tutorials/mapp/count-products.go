package mapp

func CountProducts(history [][]string) map[string]int {
	products := make(map[string]int)

	if len(history) == 0 {
		return products
	}

	for _, order := range history {
		for _, item := range order {
			products[item]++
		}
	}
	return products
}
