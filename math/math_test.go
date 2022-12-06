package math

import (
	"fmt"
	"testing"
)

type testpair struct {
	values  []float64 // input 
	average float64 // output
}

var tests = []testpair{ // list of structs (pairs)
	{[]float64{},0},
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},

}

func TestAverage(t *testing.T) {
	for i, parameter := range tests {
		fmt.Println(i)
		funcAverage, _ := Average(&parameter.values)

		if funcAverage != parameter.average {
			t.Error(
				"For these", parameter.values,
				"expected", parameter.average,
				"function returned", funcAverage,
			)
		}
	}
}
	

