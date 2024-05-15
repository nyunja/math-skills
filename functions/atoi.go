package functions

import "os"

// Convert string to integer
func Atoi(s string) int {
	result := 0
	sign := 1
	count := 0
	signIndex := 0
	for i, n := range s {
		if n == '-' {
			count++
			if count > 1 {
				os.Stdout.WriteString("Error: invalid number")
				os.Exit(0)
			} else {
				signIndex = i
				sign = -1
			}
		}
		if n == '+' {
			count++
			if count > 1 {
				os.Stdout.WriteString("Error: invalid number")
				os.Exit(0)
			} else {
				signIndex = i
				sign = 1
			}
		}
		if len(s)-signIndex < 15 {
			if n >= '0' && n <= '9' {
				result = result*10 + int(n-'0')
			}
		} else {
			sign = -1
			result = 9223372036854775807
		}
	}
	return result * sign
}
