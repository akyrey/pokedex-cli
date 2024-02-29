package pokeapi

import (
	"fmt"
	"math/rand"
	"os"
)

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show this help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokédex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Retrieve Pokémons from a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokémon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "See details about a caught Pokémon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokémons",
			callback:    commandPokedex,
		},
	}
}

func commandHelp(cfg *Config, args []string) error {
	fmt.Printf("\nWelcome to the Aky's Pokedex!\nUsage:\n")
	for _, command := range getCommands() {
		fmt.Printf("\t- %s: %s\n", command.name, command.description)
	}
	fmt.Printf("\n")
	return nil
}

func commandExit(cfg *Config, args []string) error {
	os.Exit(0)
	return nil
}

func commandMapf(cfg *Config, args []string) error {
	p, err := cfg.Client.ListLocations(cfg.NextLocations)
	if err != nil {
		return err
	}

	cfg.PreviousLocations = p.Previous
	cfg.NextLocations = p.Next

	for _, area := range p.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(cfg *Config, args []string) error {
	if cfg.PreviousLocations == nil {
		return fmt.Errorf("no previous page")
	}

	p, err := cfg.Client.ListLocations(cfg.PreviousLocations)
	if err != nil {
		return err
	}

	cfg.PreviousLocations = p.Previous
	cfg.NextLocations = p.Next

	for _, area := range p.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandExplore(cfg *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing location name")
	}

	area := args[0]
	fmt.Printf("Exploring %s...\n", area)
	explore, err := cfg.Client.ExploreLocation(area)
	if err != nil {
		return err
	}

	fmt.Printf("Found Pokémons:\n")
	for _, encounter := range explore.Encounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing Pokémon name")
	}

	name := args[0]
	fmt.Printf("Throwing a pokéball at %s...\n", name)

	pokemon, err := cfg.Client.GetPokemon(name)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	if res > 40 {
		fmt.Printf("%s ran away!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.Pokedex[pokemon.Name] = *pokemon

	return nil
}

func commandInspect(cfg *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing Pokémon name")
	}

	name := args[0]
	pokemon, ok := cfg.Pokedex[name]
	if !ok {
		return fmt.Errorf("you have not caught that Pokémon yet")
	}

	fmt.Println(pokemon)
	return nil
}

func commandPokedex(cfg *Config, args []string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("You haven't caught any Pokémon yet")
		return nil
	}

	fmt.Println("Your Pokédex:")
	for name := range cfg.Pokedex {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
