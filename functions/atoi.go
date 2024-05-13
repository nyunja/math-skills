package functions

import "os"

func Atoi(s string) int {
	result := 0
	sign := 1
	count := 0
	for _, n := range s {
		if n == '-' {
			count++
			if count > 1 {
				os.Stdout.WriteString("Error: invalid number")
				os.Exit(0)
			} else {
				sign = -1
			}
		}
		if n == '+' {
			count++
			if count > 1 {
				os.Stdout.WriteString("Error: invalid number")
				os.Exit(0)
			} else {
				sign = 1
			}
		}
		if n >= '0' && n <= '9' {
			result = result*10 + int(n-'0')
		}
	}
	return result * sign
}
