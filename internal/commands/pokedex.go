package commands

import "github.com/zulkou/pokedex/internal/api"

type Pokedex struct {
    CaughtPokemon   map[string]api.PokemonDetails
}

func NewPokedex() *Pokedex {
    return &Pokedex{
        CaughtPokemon: make(map[string]api.PokemonDetails),
    }
}
