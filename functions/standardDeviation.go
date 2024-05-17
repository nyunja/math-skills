package functions

// Calculate the standard deviation of number
func StandardDeviation(variance float64) float64 {
	var guess float64 = variance / 2.0
	for i := 0; i < 10; i++ {
		guess = (guess + variance/guess) / 2.0
	}
	// fmt.Printf("stdDeviation: %0.f\n", guess)
	return guess
}
