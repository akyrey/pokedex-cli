package main

import (
	"time"

	"github.com/akyrey/pokedex-cli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewHttpClient(10*time.Minute, 5*time.Second)
	cfg := &pokeapi.Config{
		Client:  &client,
		Pokedex: make(map[string]pokeapi.Pokemon),
	}

	pokeapi.StartRepl(cfg)
}
