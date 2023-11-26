package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func understandTypeAssertions() {
	// could be also of type "any" since it's an alias for interface{}
	var something interface{} = "hello"

	// this type assertion extracts the value as a string
	someString := something.(string)
	fmt.Println(someString)

	// this type assertion extracts the value as an int
	// this will not panic because it returns a boolean as the second value to indicate if the assertion was successful
	// when the assertion fails, the first value is the zero value for the type
	someInt, ok := something.(int)
	fmt.Println(someInt, ok)

	// this type assertion also extracts the value as an int
	// this will panic since the value is not an int
	// someInt = something.(int)
	// fmt.Println(someInt)

	// this type assertion extracts the value as a Person (struct)
	person, ok := something.(Person)
	fmt.Println(person, ok)
}
