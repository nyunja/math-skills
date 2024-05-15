package functions

// Calculate the variance of number
func Variance(a []float64, mean float64) float64 {
	var total float64
	var result float64
	for _, n := range a {
		num := n - mean
		numSq := num * num
		// fmt.Printf("squares: %0.f\n", numSq)
		total += numSq
	}
	result = total / float64(len(a))
	// fmt.Printf("variance: %0.f\n", result)
	return result
}
