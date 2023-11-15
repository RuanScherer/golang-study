package main

import (
	"fmt"
	"log"
)

func CrossPackageFunction() {
	log.Print("this function can be called from other packages because it's exported (function name starts with upper case)")
}

func packagePrivateFunction() {
	log.Print("this function can be called only from inside it's own package because it's not exported (function name starts with lower case)")
}

func myVariadicFunction(nums ...int) {
	fmt.Println(nums, " ")

	total := 0
	for _, n := range nums {
		total += n
	}
	fmt.Println("total: ", total)
}

// closure
func intSequence() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
