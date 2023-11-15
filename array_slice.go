package main

import "fmt"

func understandArray() {
	var numbers [6]bool // declare array with 6 positions, filled with zero-value
	fmt.Println(numbers)

	numbers[0] = true
	fmt.Println(numbers)
}

func understandSlice() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println(numbers)

	_ = append(numbers, 5) // doesn't take effect, need to re-assign
	fmt.Println(numbers)

	numbers = append(numbers, 5)
	fmt.Println(numbers)

	numbers = numbers[0:2] // slice the slice from 0 to 2 (exluding the 2 position)
	fmt.Println(numbers)
	fmt.Printf("slice length is %d and slice capacity is %d\n", len(numbers), cap(numbers))

	var names []string // zero-value for slice is nil
	fmt.Println(names, len(names), cap(names))
	if names == nil {
		fmt.Println("names slice is nil")
	}
}
