package main

import "log"

func understandPanic() {
	defer log.Print("finished execution") // deferred functions are called normally even when panic is invoked
	intermediate()
	log.Print("continuing after panic")
}

func intermediate() {
	defer func() {
		if r := recover(); r != nil {
			log.Print("recovered by intermediate")
		}
	}()

	log.Print("starting intermediate")
	panicGenerator()
	log.Print("finished intermediate normally")
}

func panicGenerator() {
	log.Print("starting something")
	panic("ah... something went wrong")
	log.Print("finished something normally")
}
