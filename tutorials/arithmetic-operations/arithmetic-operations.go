package arithmeticoperations

func CalculateProgress(done, total int) float64 {
	var result float64
	x := float64(done)
	y := float64(total)
	result = x / y
	return result
}
