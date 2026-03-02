package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(c *config) {

	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scan.Scan()
		c.cmd = cleanInput(scan.Text())
		if len(c.cmd) == 0 {
			continue
		}
		cmd := c.cmd[0]
		value, exist := getCommand()[cmd]
		if exist {
			err := value.callback(c)
			if err != nil {
				fmt.Println(err)
			}
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
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location-area>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "Catch <pokemon>",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "Inspect <pokemon>",
			description: "Inspect a pokemon that was caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Inspect the pokedex",
			callback:    commandPokedex,
		},
	}
}
