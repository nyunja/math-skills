package functions

// Convert integer to string
func Itoa(n int) string {
	result := ""
	sign := ""
	if n < 0 {
		sign = "-"
		n = n * -1
	}
	for i := 0; n/10 > 0; i++ {
		result = string(rune(n%10+'0')) + result
		n /= 10
	}
	if n/10 == 0 {
		result = string(rune(n+'0')) + result
	}
	return sign + result
}
