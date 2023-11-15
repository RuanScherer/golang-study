package main

import "fmt"

var person map[string]any

func understandMaps() {
	fmt.Println(person)
	if person == nil {
		fmt.Println("person is nil")
	}

	person = make(map[string]any)
	person["name"] = "Ruan"
	fmt.Println(person)

	person = map[string]any{
		"name": "Vitor",
		"age":  18,
	}
	fmt.Println(person)

	fmt.Println("age is", person["age"])

	delete(person, "age")

	fmt.Println("age is now", person["age"])

	// check if key is present
	value, ok := person["age"]
	if ok {
		fmt.Println("age is", value)
	} else {
		fmt.Println("age is not set")
	}
}
