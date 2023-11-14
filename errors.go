package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func gettingError() {
	_, err := os.Open("fake-file.txt")
	if err != nil {
		log.Fatal("error reading file")
	}
	log.Print("finished") // is not executed
}

func buildingErrors() {
	divide := func(number1, number2 float32) (float32, error) {
		if number2 == 0 {
			return 0, errors.New("math: cannot divide by zero")
		}
		return number1 / number2, nil
	}

	result, err := divide(2, 0)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print(result)
}

type ZeroDivisionError float32

func (z ZeroDivisionError) Error() string {
	return fmt.Sprintf("math: cannot divide by %g", z)
}

func buildingCustomErrors() {
	divide := func(number1, number2 float32) (float32, error) {
		if number2 == 0 {
			return 0, ZeroDivisionError(number2)
		}
		return number1 / number2, nil
	}

	result, err := divide(2, 0)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print(result)
}
