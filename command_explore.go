package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	pokecache "github.com/Util787/pokedex/internal"
)

type FoundedPokemons struct {
	Pokemon_encounters []struct {
		Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name string `json:"name"`
}

func commandExplore(a *config, cch *pokecache.Cache, areaname string) error {

	if areaname == "" {
		log.Println("Given zerovalue")
		return errors.New("Given zerovalue")
	}

	explurl := "https://pokeapi.co/api/v2/location-area/" + areaname

	var js []byte

	if val, ok := cch.Get(explurl); ok {
		js = val
	} else {
		resp, err := http.Get(explurl)
		if err != nil {
			return err
		}

		js, err = io.ReadAll(resp.Body) // if you write := here you'll get pokemons wrong because of shadowing
		if err != nil {
			return err
		}

		cch.Add(explurl, js)

		defer resp.Body.Close()
	}

	data := FoundedPokemons{}

	err := json.Unmarshal(js, &data)
	if err != nil {
		return err
	}

	for _, pokemons := range data.Pokemon_encounters {
		fmt.Println(" - " + pokemons.Name)
	}

	return nil
}
