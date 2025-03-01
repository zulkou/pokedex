package commands

import (
	"fmt"
	"math"
	"math/rand"
	"os"

	"github.com/zulkou/pokedex/internal/api"
	"github.com/zulkou/pokedex/internal/config"
	"github.com/zulkou/pokedex/internal/pokecache"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(cfg *config.Config, cache *pokecache.Cache, pokedex *Pokedex, args ...string) error
}

var Commands map[string]CliCommand

func InitializeCommand() {
    Commands = map[string]CliCommand{
        "exit": {
            Name: "exit",
            Description: "Exit the Pokedex",
            Callback: commandExit,
        },
        "help": {
            Name: "help",
            Description: "Displays a help message",
            Callback: commandHelp,
        },
        "map": {
            Name: "map",
            Description: "Displays 20 locations in Pokemon World",
            Callback: commandMap,
        },
        "mapb": {
            Name: "mapb",
            Description: "Return to last 20 locations showed",
            Callback: commandMapBack,
        },
        "explore": {
            Name: "explore <location-area>",
            Description: "Explore given location area",
            Callback: commandExplore,
        },
        "catch": {
            Name: "catch <pokemon>",
            Description: "Try catch given pokemon",
            Callback: commandCatch,
        },
        "inspect": {
            Name: "inspect <pokemon>",
            Description: "Inspect caught pokemon",
            Callback: commandInspect,
        },
    }
}

func commandExit(cfg *config.Config, cache *pokecache.Cache, pokedex *Pokedex, args ...string) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp(cfg *config.Config, cache *pokecache.Cache, pokedex *Pokedex, args ...string) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
    for _, v := range Commands {
        fmt.Printf("%v: %v\n", v.Name, v.Description)
    }
    return nil
}

func commandMap(cfg *config.Config, cache *pokecache.Cache, pokedex *Pokedex, args ...string) error {
    urlToUse := cfg.BaseURL
    if cfg.NextPageURL != nil {
        urlToUse = *cfg.NextPageURL
    }

    data, err := api.FetchLocation(urlToUse, cache)
    if err != nil {
        return fmt.Errorf("failed to fetch locations: %w", err)
    }

    for _, loc := range data.Results {
        fmt.Println(loc.Name)
    }
    cfg.NextPageURL = data.Next
    cfg.PreviousPageURL = data.Previous

    return nil
}

func commandMapBack(cfg *config.Config, cache *pokecache.Cache, pokedex *Pokedex, args ...string) error {
    if cfg.PreviousPageURL == nil {
        fmt.Println("You're on the first page, there's no going back!")
        return nil
    }

    data, err := api.FetchLocation(*cfg.PreviousPageURL, cache)
    if err != nil {
        return fmt.Errorf("failed to fetch locations: %w", err)
    }

    for _, loc := range data.Results {
        fmt.Println(loc.Name)
    }
    cfg.NextPageURL = data.Next
    cfg.PreviousPageURL = data.Previous

    return nil
}

func commandExplore(cfg *config.Config, cache *pokecache.Cache, pokedex *Pokedex, args ...string) error {
    currAreaURL := cfg.LocationAreaURL + args[0]

    data, err := api.FetchExplore(currAreaURL, cache)
    if err != nil {
        return fmt.Errorf("failed to fetch exploration area: %w", err)
    }

    fmt.Printf("Exploring %v...\n", args[0])
    fmt.Println("Found Pokemon:")
    for _, encounter := range data.Encounters{
        fmt.Println("-", encounter.Pokemon.Name)
    }

    return nil
}

func commandCatch(cfg *config.Config, cache *pokecache.Cache, pokedex *Pokedex, args ...string) error {
    pokemonURL := cfg.PokemonURL + args[0]

    pokemon, err := api.FetchPokemon(pokemonURL, cache)
    if err != nil {
        return fmt.Errorf("failed to fetch pokemon data: %w", err)
    }

    fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)

    success := rand.Float64() < math.Exp(-((float64(pokemon.Chance) - 100) / 71.083) - math.Ln2)
    if success {
        fmt.Printf("%v was caught!\n", pokemon.Name)
        pokedex.CaughtPokemon[pokemon.Name] = *pokemon
    } else {
        fmt.Printf("%v escaped!\n", pokemon.Name)
    }

    return nil
}

func commandInspect(cfg *config.Config, cache *pokecache.Cache, pokedex *Pokedex, args ...string) error {
    pokemon, exists := pokedex.CaughtPokemon[args[0]]; if !exists {
        fmt.Println("you have not caught that pokemon")

        return nil
    }

    fmt.Printf("Name: %v\n", pokemon.Name)
    fmt.Printf("Height: %v\n", pokemon.Height)
    fmt.Printf("Weight: %v\n", pokemon.Weight)
    fmt.Println("Stats:")
    for key, val := range pokemon.Stats {
        fmt.Printf(" -%v: %v\n", key, val)
    }
    fmt.Println("Types:")
    for _, t := range pokemon.Types {
        fmt.Printf(" - %v\n", t)
    }
    
    return nil
}
