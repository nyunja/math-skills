package functions

import "os"

// Convert string values to float64
func ToFloat(a []string) []float64 {
	var result []float64
	for _, num := range a {
		if float64(Atoi(num)) <= -9223372036854775807 || float64(Atoi(num)) >= 9223372036854775807 {
			os.Stdout.WriteString("Error: overflow at value " + num + "\n")
			os.Exit(0)
		} else {
			result = append(result, float64(Atoi(num)))
		}
	}
	return result
}
