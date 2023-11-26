package main

import "fmt"

type message string

func (m message) Read() string {
	return string(m)
}

func understandCustomTypes() {
	var msg message = "hello"
	fmt.Printf("%T: %v\n", msg, msg)

	// this will not compile because the underlying type of msg is string
	// although msg is based on the type string, it is not the same type
	// var invalidMsg string = msg

	// custom types can also have methods
	fmt.Println(msg.Read())
}
