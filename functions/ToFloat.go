package functions

// Convert string values to float64
func ToFloat(a []string) []float64 {
	var result []float64
	for _, num := range a {
		if float64(Atoi(num)) == -9223372036854775807 {
			continue
		} else {
			result = append(result, float64(Atoi(num)))
		}
	}
	return result
}
