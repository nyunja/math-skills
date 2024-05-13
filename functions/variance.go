package functions

func Variance(a []float64, mean float64) float64 {
	var total float64
	var result float64
	for _, n := range a {
		num := n - mean
		num = num * num
		total += num
	}
	result = total / float64(len(a))
	return result
}
