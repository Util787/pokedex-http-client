package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	pokecache "github.com/Util787/pokedex/internal"
)

// RespShallowLocations -
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMapf(conf *config, cch *pokecache.Cache, s string) error {
	var body []byte

	if val, ok := cch.Get(*conf.nextLocationsURL); ok {
		body = val
	} else {

		resp, err := http.Get(*conf.nextLocationsURL)
		if err != nil {
			return err
		}

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		cch.Add(*conf.nextLocationsURL, body)

		defer resp.Body.Close()
	}
	data := RespShallowLocations{}

	err := json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}

	conf.nextLocationsURL = data.Next
	conf.prevLocationsURL = data.Previous

	return nil
}

func commandMapb(conf *config, cch *pokecache.Cache, s string) error {
	if conf.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return errors.New("you're on the first page")
	}
	var body []byte
	if val, ok := cch.Get(*conf.prevLocationsURL); ok {
		body = val
	} else {
		resp, err := http.Get(*conf.prevLocationsURL)
		if err != nil {
			return err
		}

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		cch.Add(*conf.prevLocationsURL, body)

		defer resp.Body.Close()
	}
	data := RespShallowLocations{}

	err := json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}

	conf.nextLocationsURL = data.Next
	conf.prevLocationsURL = data.Previous

	return nil
}
