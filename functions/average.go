package functions

// Calculate the average of a list of numbers
func Average(numbers []float64) float64 {
	var result float64
	for _, n := range numbers {
		result += n
	}
	average := result / float64(len(numbers))
	return average
}
