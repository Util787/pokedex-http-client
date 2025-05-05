package main

import (
	"fmt"

	pokecache "github.com/Util787/pokedex/internal"
)

func commandHelp(*config, *pokecache.Cache, string) error {
	fmt.Println("Welcome to the Pokedex! \n Usage:")
	for _, vals := range getCommands() {
		fmt.Printf("%s: %s\n", vals.name, vals.description)
	}

	return nil
}
