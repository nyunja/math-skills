package functions

import (
	"fmt"
	"os"
)

// Print the solutions to the terminal
func HandleCalculation(numStr []string) {
	numStr = SortList(numStr)
	numbers := ToFloat(numStr)
	// fmt.Println(numbers[0] * 10)

	mean := Average(numbers)
	median := Median(numbers)
	variance := Variance(numbers, mean)
	stdDev := StandardDeviation(variance)
	//os.Stdout.WriteString("Average: ")
	fmt.Printf("Average: %0.f\n", mean)
	//os.Stdout.WriteString(Itoa(round(mean)) + "\n")
	os.Stdout.WriteString("Median: ")
	os.Stdout.WriteString(Itoa(round(median)) + "\n")
	fmt.Printf("Variance: %0.f\n", float64(variance))
	//os.Stdout.WriteString(Itoa(round(variance)))
	//os.Stdout.WriteString("Standard Deviation: ")
	fmt.Printf("Standard Deviation: %0.f\n", float64(stdDev))
	//os.Stdout.WriteString(Itoa(round(stdDev)) + "\n")
}
