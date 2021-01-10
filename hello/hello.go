package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties for logger
	// log entry prefix + a flag to disable printing the time, source file and line #
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message + print
	message, err := greetings.Hello("BenTheDev")
	message1, err := greetings.Hello("")

	// if an error was returned, print it in the console + exit
	if err != nil {
		log.Fatal(err)
	}

	// print message if no error was returned
	fmt.Println(message)
	fmt.Println(message1)
}
