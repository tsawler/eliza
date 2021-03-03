package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(fmt.Sprintf("%s %s %s", "\033[1m", Intro(), "\033[0m"))

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "quit" {
			break
		} else {
			resp := fmt.Sprintf("%s%s%s", "\033[1m", Response(userInput), "\033[0m")
			fmt.Println(resp)
		}
	}
}
