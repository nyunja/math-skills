package functions

func Median(a []float64) float64 {
	var median float64
	if len(a)%2 == 0 {
		index := (len(a) / 2) - 1
		num1 := a[index]
		num2 := a[index+1]
		median = (float64(num1) + float64(num2)) / 2
	} else {
		index := (len(a) / 2) - 1
		median = a[index+1]
	}
	return median
}
