package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Craete bufio.Scanner reading from os.Stdin
	scanner := bufio.NewScanner(os.Stdin)

	//infinite for loop; execute once for every command user inputs
	for {
		fmt.Print("Pokedex > ")

		for scanner.Scan() {
			input := scanner.Text()
			commands := getCommands()
			//check input against supportedCommands map (found in repl.go)
			if _, ok := commands[input]; ok {
				commands[input].callback()
			} else {
				fmt.Println("Unknown command")
			}

			fmt.Print("Pokedex > ")
		}
		//check for scanner error
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading user input: ", err)
		}
	}
}
