package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	pokecache "github.com/Util787/pokedex/internal"
)

var PokeDex map[string]int

type PokeExp struct {
	BaseExp int `json:"base_experience"`
}

func commandCatch(conf *config, cch *pokecache.Cache, pokemonname string) error {
	fmt.Printf("Throwing a Pokeball at %v... \n", pokemonname)

	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + pokemonname)
	if err != nil {
		return err
	}

	js, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// can do a cacheentry here

	defer resp.Body.Close()

	data := PokeExp{}
	err = json.Unmarshal(js, &data)
	if err != nil {
		return err
	}

	EscapeChance := data.BaseExp / 10

	PlayerChance := rand.Intn(101)

	if PlayerChance >= EscapeChance {
		fmt.Println(pokemonname + " was caught!")
		if PokeDex == nil {
			PokeDex = make(map[string]int)
		}
		PokeDex[pokemonname]++
	} else {
		fmt.Println(pokemonname + " escaped")
	}
	fmt.Println(`-------------------------------------------------------
Your Pokemons:`)
	for k, v := range PokeDex {
		fmt.Println(k, ":", v)
	}
	return nil

}
