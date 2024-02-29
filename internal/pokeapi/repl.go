package pokeapi

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/akyrey/pokedex-cli/internal/pokecache"
)

type Config struct {
	Cache             *pokecache.Cache
	Client            *Client
	NextLocations     *string
	PreviousLocations *string
	Pokedex           map[string]Pokemon
}

type cliCommand struct {
	callback    func(*Config, []string) error
	name        string
	description string
}

func StartRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Aky's Pokédex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Command not found")
			continue
		}

		err := command.callback(cfg, words[1:])
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
