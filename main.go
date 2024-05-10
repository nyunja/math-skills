package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Stdout.WriteString("Error: please provide the required arguments")
		return
	}
	fileName := os.Args[1]
	if fileName != "data.txt" {
		os.Stdout.WriteString("Error: please provide the correct file")
		return
	}
	content, err := os.ReadFile(fileName)
	if err != nil {
		os.Stdout.WriteString("Error: unable to read " + fileName + " content")
		return
	}
	if len(content) == 0 {
		os.Stdout.WriteString("Error: " + fileName + " is empty")
		return
	}
	numbers := splitString(string(content), "\n")
	numbers = sortList(numbers)
	mean := Average(numbers)
	variance := Variance(numbers, mean)
	stdDev := standardDev(variance)
	fmt.Println(stdDev)
	fmt.Println(variance)
	handleStats(numbers)
}

func splitString(s string, sep string) []string {
	result := []string{}
	token := ""
	for i := 0; i < len(s)-len(sep); i++ {
		if s[i:len(sep)] == sep {
			result = append(result, token)
			token = ""
		} else {
			token += string(s[i])
		}
	}
	result = append(result, token)
	token = ""
	return result
}

func handleStats(a []string) {
}

func Average(numbers []string) float64 {
	var result float64
	for _, n := range numbers {
		result += float64(Atoi(n))
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

func Median(a []string) float64 {
	var median float64
	if len(a)%2 == 0 {
		index := (len(a) / 2) - 1
		num1 := Atoi(a[index])
		num2 := Atoi(a[index+1])
		median = (float64(num1) + float64(num2)) / 2
	} else {
		index := (len(a) / 2) - 1
		median = float64(Atoi(a[index+1]))
	}
	return median
}

func Variance(a []string, mean float64) float64 {
	var total float64
	var result float64
	for _, n := range a {
		num := float64(Atoi(n))
		num = num - mean
		num = num * num
		total += num
	}
	result = total / float64(len(a))
	return result
}

func standardDev(variance float64) float64 {
	// n := int(variance)
	var result float64
	for i := 0.0; i*i == variance; i += 0.1 {
		if i*i == variance {
			result = i
		}
	}
	return result
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
	if n == 0 {
		result = string(rune(n+'0')) + result
	}
	return sign + result
}
