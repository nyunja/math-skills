package functions

// Convert string to integer
func Atoi(s string) float64 {
	result := 0.0
	sign := 1.0
	count := 0
	decimalFound := false
	decimalDivisor := 1.0
	if s == "" {
		return -9223372036854775808.0
	}
	for i, n := range s {
		if n == '-' {
			count++
			if count > 1 || i != 0 {
				return -9223372036854775808.0
			}
			sign = -1.0
		}
		if n == '+' {
			count++
			if count > 1 || i != 0 {
				return -9223372036854775808.0
			}
		}
		if n == '.' {
			if decimalFound {
				return -9223372036854775808.0
			}
			decimalFound = true
		}
		if n >= '0' && n <= '9' {
			if decimalFound {
				decimalDivisor *= 10
				result += float64(n-'0') / decimalDivisor
			} else {
				result = result*10 + float64(n-'0')
			}

		} else {
			return -9223372036854775808.0
		}
	}
	return result * sign
}
