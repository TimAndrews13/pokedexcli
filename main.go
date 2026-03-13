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
	pokedex := make(map[string]Pokemon)
	cfg := config{
		Cache:   cache,
		Pokedex: pokedex,
	}

	//infinite for loop; execute once for every command user inputs
	for {
		fmt.Print("Pokedex > ")

		for scanner.Scan() {
			input := scanner.Text()
			words := cleanInput(input)
			command := words[0]
			arg := ""
			if len(words) > 1 {
				arg = words[1]
			}

			commands := getCommands()
			//check input against supportedCommands map (found in repl.go)
			if _, ok := commands[command]; ok {
				commands[command].callback(&cfg, arg)
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
