package functions

// Convert string values to float64
func ToFloat(a []string) []float64 {
	var result []float64
	for _, num := range a {
		if Atoi(num) == -9223372036854775808.0 || Atoi(num) > 9223372036854775807.0 {
			continue
		} else {
			result = append(result, float64(Atoi(num)))
		}
	}
	return result
}
