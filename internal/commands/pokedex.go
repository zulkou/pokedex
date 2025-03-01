package commands

import "github.com/zulkou/pokedex/internal/api"

type Pokedex struct {
    CaughtPokemon   map[string]api.Pokemon
}

func NewPokedex() *Pokedex {
    return &Pokedex{
        CaughtPokemon: make(map[string]api.Pokemon),
    }
}
