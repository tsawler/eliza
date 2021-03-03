package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// reader is a variable used to read whatever the user types in on the console
	// reader requires some input source (anything that satisfies the io.Reader interface),
	// and in this case we'll use "standard in" (the keyboard/console).
	reader := bufio.NewReader(os.Stdin)

	// Print the welcome message and instructions to the screen. Note that we used a function as a parameter
	// to another function. Go has first class functions.
	fmt.Println(Intro())

	// Execute everything in this loop until the user enters 'quit' and presses return.
	for {
		// Print out a prompt
		fmt.Print("-> ")

		// Wait for the user to type something and press enter
		userInput, _ := reader.ReadString('\n')

		// Strip the carriage return off of whatever they typed
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			// end program
			break
		} else {
			// go build the response and write it to the console
			fmt.Println(Response(userInput))
		}
	}
}
