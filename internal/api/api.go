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
    }                   `json:"results"`
    Next        *string `json:"next"`
    Previous    *string `json:"previous"`
}

type ExploreArea struct {
    Encounters []struct {
        Pokemon struct {
            Name    string  `json:"name"`
        }                   `json:"pokemon"`
    }                       `json:"pokemon_encounters"`
}

type PokemonDetails struct {
    Name    string  `json:"name"`
    Chance  int     `json:"base_experience"`
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

func FetchPokemon(url string, cache *pokecache.Cache) (*PokemonDetails, error) {
    if data, found := cache.Get(url); found {
        var pokemon PokemonDetails
        err := json.Unmarshal(data, &pokemon)
        return &pokemon, err
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

    var result PokemonDetails
    err = json.Unmarshal(body, &result)
    if err != nil {
        return nil, fmt.Errorf("error parsing json: %w", err)
    }

    return &result, nil
}
