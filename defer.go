package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func usingDefer() error {
	source, err := os.Open("my-file.txt")
	if err != nil {
		return errors.New("failed to open source file")
	}
	defer source.Close()

	destination, err := os.Create("my-file-copy.txt")
	if err != nil {
		return errors.New("failed to create copy file")
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return errors.New("failed to copy content from source file")
	}

	return nil
}

func understandDeferArguments() {
	i := 0

	// it's going to print 0 since arguments are evaluated at the moment that function is deferred
	defer fmt.Println(i)

	i++
	return
}

func understandDeferCallOrder() {
	// deferred functions are called in the "Last in First out" order
	for i := 0; i < 4; i++ {
		defer fmt.Println(i) // example from golang blog
	}
}
