package main

import (
	"time"

	pokecache "github.com/Util787/pokedex/internal"
)

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
}

func main() {
	baseurl := "https://pokeapi.co/api/v2/location-area"
	cfg := &config{
		nextLocationsURL: &baseurl,
	}
	cch := pokecache.NewCache(5 * time.Second)
	start(cfg, cch)
}
