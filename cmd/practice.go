package main

import (
	"fmt"

	"github.com/Vfulgham/go-mux-api/math"
)

func add(args ...int) int {
	var total int
	for _, num := range args {
		total += num
	}

	return total
}

func practice() {

	total := add(1, 2, 3)
	fmt.Println(total)

	var message string = "Hi"
	fmt.Printf("%v there \n", message)

	mySlice := []string{"Hi", "there", "my", "name", "is", "Vicki"}
	for i, name := range mySlice {
		fmt.Print(i, "-", name, ", ")
	}

	//const vic, steve, kylis = "forty", "forty2", "forty3"

	/*
		const (
			vic2   = "forty"
			steve2 = "forty2"
			kylis2 = "forty3"
		)
	*/

	var myArr [2]int
	myArr[0] = 1
	myArr[1] = 2

	//var myArr2 = [2]int{1,2}

	//var mySlice2 = []float64{}

	var mySlice3 = make([]float64, 0)
	mySlice4 := append(mySlice3, 2)
	fmt.Println(mySlice4)

	// loop
	for i := 0; i < 3; i++ {
		fmt.Printf("Hi there %v \n", i)
	}

	// map 1
	var x = make(map[string]int)
	x["one"] = 1
	x["two"] = 2
	fmt.Println(x)
	delete(x, "two")
	fmt.Println(x)

	// map 2
	var y = map[string]int{
		"one": 1,
		"two": 2,
	}

	for x, y := range y {
		fmt.Println(x, y)
	}

	// ------------------------------------------------------
	xs := []float64{}
	avg, _ := math.Average(&xs)
	fmt.Println(avg)

	var nil_slice []int
	var empty_slice = []int{}
	var emptySlice1 = make([]int, 0)

	fmt.Println(nil_slice == nil, len(nil_slice), cap(nil_slice))
	fmt.Println(empty_slice == nil, len(empty_slice), cap(empty_slice))
	fmt.Println(emptySlice1 == nil, len(empty_slice), cap(empty_slice))

}
