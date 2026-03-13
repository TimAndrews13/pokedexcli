package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/timandrews/pokedexcli/internal/pokecache"
)

func main() {
	//Craete bufio.Scanner reading from os.Stdin
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(5 * time.Second)
	cfg := config{
		Cache: cache,
	}

	//infinite for loop; execute once for every command user inputs
	for {
		fmt.Print("Pokedex > ")

		for scanner.Scan() {
			input := scanner.Text()
			commands := getCommands()
			//check input against supportedCommands map (found in repl.go)
			if _, ok := commands[input]; ok {
				commands[input].callback(&cfg)
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
