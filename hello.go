package main

import (
	"fmt"
	"log"
	"os"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	if len(os.Args) >= 2 {
		log.Fatal("Dont forget the args")
		return
	}

	// Get a greeting message and print it.
	nome := os.Args[1] + " " + os.Args[2]
	message, err := greetings.Hello(nome + " do canal " + nome)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
