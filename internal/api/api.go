package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/zulkou/pokedex/internal/pokecache"
)

type LocationArea struct {
    Results []struct {
        Name    string  `json:"name"`
    } `json:"results"`
    Next        *string `json:"next"`
    Previous    *string `json:"previous"`
}

type ExploreArea struct {
    Encounters []struct {
        Pokemon struct {
            Name    string  `json:"name"`
        } `json:"pokemon"`
    } `json:"pokemon_encounters"`
}

type Pokemon struct {
    Name    string
    Chance  int
    Height  int
    Weight  int
    Stats   map[string]int
    Types   []string
}

type RawPokemon struct {
    Name    string              `json:"name"`
    Chance  int                 `json:"base_experience"`
    Height  int                 `json:"height"`
    Weight  int                 `json:"weight"`
    Stats   []struct {
        BaseStat    int         `json:"base_stat"`
        Stat        struct {
            Name    string      `json:"name"`
        } `json:"stat"`
    } `json:"stats"`
    Types   []struct {
        Type        struct {
            Name    string      `json:"name"`
        } `json:"type"`
    } `json:"types"`
}

func FetchLocation(url string, cache *pokecache.Cache) (*LocationArea, error) {
    if data, found := cache.Get(url); found {
        var location LocationArea
        err := json.Unmarshal(data, &location)
        return &location, err
    }

    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error making API request: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err!= nil {
        return nil, fmt.Errorf("error reading response body: %w", err)
    }

    cache.Add(url, body)

    var result LocationArea
    err = json.Unmarshal(body, &result)
    if err != nil {
        return nil, fmt.Errorf("error parsing json: %w", err)
    }

    return &result, nil
}

func FetchExplore(url string, cache *pokecache.Cache) (*ExploreArea, error) {
    if data, found := cache.Get(url); found {
        var explore ExploreArea
        err := json.Unmarshal(data, &explore)
        return &explore, err
    }

    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error making API request: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err!= nil {
        return nil, fmt.Errorf("error reading response body: %w", err)
    }

    cache.Add(url, body)

    var result ExploreArea
    err = json.Unmarshal(body, &result)
    if err != nil {
        return nil, fmt.Errorf("error parsing json: %w", err)
    }

    return &result, nil
}

func FetchPokemon(url string, cache *pokecache.Cache) (*Pokemon, error) {
    if data, found := cache.Get(url); found {
        var pokemon RawPokemon
        err := json.Unmarshal(data, &pokemon)
        return pokemon.pokemonTransformer(), err
    }

    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error making API request: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %w", err)
    }

    cache.Add(url, body)

    var result RawPokemon
    err = json.Unmarshal(body, &result)
    if err != nil {
        return nil, fmt.Errorf("error parsing json: %w", err)
    }

    return result.pokemonTransformer(), nil
}

func (raw *RawPokemon)pokemonTransformer() (*Pokemon) {
    pokemon := &Pokemon {
        Name: raw.Name,
        Chance: raw.Chance,
        Weight: raw.Weight,
        Height: raw.Height,
        Stats: make(map[string]int),
        Types: make([]string, 0),
    }

    for _, statObj := range raw.Stats {
        pokemon.Stats[statObj.Stat.Name] = statObj.BaseStat
    }
    
    for _, typeObj := range raw.Types {
        pokemon.Types = append(pokemon.Types, typeObj.Type.Name)
    }

    return pokemon
}
