package main

import pokecache "github.com/Util787/pokedex/internal"

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache, string) error // should name it commandHandler
}

// returns map of registered commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Exploring <Area name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch <Pokemon name>",
			callback:    commandCatch,
		},
	}
}
