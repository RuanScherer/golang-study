package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	myVariadicFunction(numbers...)

	nextInt := intSequence()

	for i := 0; i < 3; i++ {
		fmt.Println(nextInt())
	}
}
