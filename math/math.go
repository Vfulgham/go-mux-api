package math

import "fmt"

// finds the average of a series of numbers
func Average(xs *[]float64) (float64, error) {
	if len(*xs) == 0 {
		return 0, fmt.Errorf("new error")
	}
	total := float64(0)
	for _, x := range *xs {
		total += x
	}
	return total / float64(len(*xs)), nil
}
