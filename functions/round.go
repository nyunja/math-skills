package functions

// Round off the number to the nearest integer
func round(f float64) int {
	if f < 0 {
		return int(f - 0.5)
	} else {
		return int(f + 0.5)
	}
}
