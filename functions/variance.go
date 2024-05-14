package functions

// Calculate the variance of number
func Variance(a []float64, mean float64) float64 {
	var total float64
	var result float64
	for _, n := range a {
		num := n - mean
		numSq := num * num
		total += numSq
	}
	result = total / float64(len(a))
	return result
}
