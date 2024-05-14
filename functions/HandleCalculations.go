package functions

import (
	"fmt"
	"os"
)

// Print the solutions to the terminal
func HandleCalculation(numStr []string) {
	numStr = SortList(numStr)
	numbers := ToFloat(numStr)
	fmt.Println(numbers[0]*10)

	mean := Average(numbers)
	median := Median(numbers)
	variance := Variance(numbers, mean)
	stdDev := StandardDeviation(variance)
	os.Stdout.WriteString("Average: ")
	os.Stdout.WriteString(Itoa(round(mean)) + "\n")
	os.Stdout.WriteString("Median: ")
	os.Stdout.WriteString(Itoa(round(median)) + "\n")
	os.Stdout.WriteString("Variance: ")
	os.Stdout.WriteString(Itoa(round(variance)))
	os.Stdout.WriteString("Standard Deviation: ")
	os.Stdout.WriteString(Itoa(round(stdDev)) + "\n")
}
