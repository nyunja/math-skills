package main

import (
	"os"
	"runtime"
)

func main() {
	if len(os.Args) != 2 {
		os.Stdout.WriteString("Error: please provide the required arguments\n")
		return
	}
	fileName := os.Args[1]
	if fileName != "data.txt" {
		os.Stdout.WriteString("Error: please provide the correct file\n")
		return
	}
	content, err := os.ReadFile(fileName)
	if err != nil {
		os.Stdout.WriteString("Error: unable to read " + fileName + " content\n")
		return
	}
	if len(content) == 0 {
		os.Stdout.WriteString("Error: " + fileName + " is empty\n")
		return
	}
	var numStr []string

	ops := runtime.GOOS
	switch ops {
	case "windows":
		numStr = splitString(string(content), "\r\n")
	case "linux":
		numStr = splitString(string(content), "\n")
	}

	numStr = sortList(numStr)
	numbers := toFloat(numStr)

	mean := Average(numbers)
	median := Median(numbers)
	variance := Variance(numbers, mean)
	stdDev := standardDev(variance)
	os.Stdout.WriteString("Average: ")
	os.Stdout.WriteString(Itoa(round(mean)) + "\n")
	os.Stdout.WriteString("Median: ")
	os.Stdout.WriteString(Itoa(round(median)) + "\n")
	os.Stdout.WriteString("Variance: ")
	os.Stdout.WriteString(Itoa(round(variance)) + "\n")
	os.Stdout.WriteString("Standard Deviation: ")
	os.Stdout.WriteString(Itoa(round(stdDev)) + "\n")
}

func round(f float64) int {
	if f < 0 {
		return int(f - 0.5)
	} else {
		return int(f + 0.5)
	}
}

func splitString(s string, sep string) []string {
	result := []string{}
	token := ""
	for i := 0; i < len(s); i++ {
		if i < len(s)-len(sep) && s[i:i+len(sep)] == sep {
			result = append(result, token)
			token = ""
			i = i + len(sep) - 1
		} else {
			token += string(s[i])
		}
	}
	result = append(result, token)
	return result
}

func toFloat(a []string) []float64 {
	var result []float64
	for _, num := range a {
		result = append(result, float64(Atoi(num)))
	}
	return result
}

func Average(numbers []float64) float64 {
	var result float64
	for _, n := range numbers {
		result += n
	}
	average := result / float64(len(numbers))
	return average
}

func sortList(a []string) []string {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

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

func standardDev(variance float64) float64 {
	var guess float64 = variance / 2.0
	for i := 0; i < 10; i++ {
		guess = (guess + variance/guess) / 2.0
	}
	return guess
}

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
