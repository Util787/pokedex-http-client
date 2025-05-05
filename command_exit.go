package main

import (
	"fmt"
	"os"

	pokecache "github.com/Util787/pokedex/internal"
)

func commandExit(*config, *pokecache.Cache, string) error {
	fmt.Print("Closing the Pokedex... Goodbye! \n")
	os.Exit(0)
	return nil
}
