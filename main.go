package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// reader is a variable used to read whatever the user types in on the console
	reader := bufio.NewReader(os.Stdin)

	// print the welcome message and instructions to the screen
	fmt.Println(Intro())

	// execute everything in this loop until the user enters 'quit'
	for {
		// print out a prompt
		fmt.Print("-> ")

		// wait for the user to type something and press enter
		userInput, _ := reader.ReadString('\n')

		// strip the carriage return off of whatever they typed
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
