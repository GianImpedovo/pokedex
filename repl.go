package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {

	scan := bufio.NewScanner(os.Stdin)
	c := &config{Next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20", Previous: ""}
	for {
		fmt.Print("Pokedex > ")
		scan.Scan()
		cleanWords := cleanInput(scan.Text())
		if len(cleanWords) == 0 {
			continue
		}
		cmd := cleanWords[0]
		value, exist := getCommand()[cmd]
		if exist {
			_ = value.callback(c)
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	Next     string
	Previous string
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show this help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Obtain the area lcoation on the pokeapi",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Obtain the previous area lcoation on the pokeapi",
			callback:    commandMapb,
		},
	}
}
