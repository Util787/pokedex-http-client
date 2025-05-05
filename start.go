package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokecache "github.com/Util787/pokedex/internal"
)

func start(cfg *config, cch *pokecache.Cache) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		userinput := cleanInput(reader.Text())
		if len(userinput) == 0 {
			continue
		}

		if command, exist := getCommands()[userinput[0]]; exist {
			if len(userinput) > 1 && len(userinput) < 3 && command.name == ("explore") || command.name == ("catch") {
				command.callback(cfg, cch, userinput[1])
			} else {
				command.callback(cfg, cch, "")
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	lowers := strings.ToLower(text)
	words := strings.Fields(lowers)
	
	return words
}
