package main

import (
	"math-skills/functions"
	"testing"
)

func TestAverage(t *testing.T) {
	result := functions.Average([]float64{19.0, 11.0, 15.0})
	var expected float64 = 15
	if result != expected {
		t.Errorf("Try again")
		return
	}
}

func TestMedian(t *testing.T) {
	result := functions.Median(functions.ToFloat(functions.SortList(([]string{"19", "11", "15"}))))
	var expected float64 = 15
	if result != expected {
		t.Errorf("Expected %v got %v\n", expected, result)
		t.Errorf("Try again")
		return
	}
}

func TestVariance(t *testing.T) {
	result := functions.Variance([]float64{19.0, 11.0, 15.0}, 15)
	var expected float64 = 10.666666666666666
	if result != expected {
		t.Errorf("Error: expected %v got %v\n", expected, result)
		t.Errorf("Try again")
		return
	}
}

func TestSortList(t *testing.T) {
	result := functions.SortList([]string{"19", "11", "15"})
	var expected []string = []string{"11", "15", "19"}
	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Errorf("Try again")
			return
		}
	}
}

func TestStdDev(t *testing.T) {
	result := functions.StandardDeviation(10.666666666666666)
	var expected float64 = 3.265986323710904
	if result != expected {
		t.Errorf("Error: expected %v got %v\n", expected, result)
		t.Errorf("Try again")
		return
	}
}
