package commands

import (
	"fmt"
	"os"

	"github.com/zulkou/pokedex/internal/api"
	"github.com/zulkou/pokedex/internal/config"
	"github.com/zulkou/pokedex/internal/pokecache"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(cfg *config.Config, cache *pokecache.Cache, args ...string) error
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
            Name: "explore",
            Description: "Explore given location area",
            Callback: commandExplore,
        },
    }
}

func commandExit(cfg *config.Config, cache *pokecache.Cache, args ...string) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp(cfg *config.Config, cache *pokecache.Cache, args ...string) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
    for _, v := range Commands {
        fmt.Printf("%v: %v\n", v.Name, v.Description)
    }
    return nil
}

func commandMap(cfg *config.Config, cache *pokecache.Cache, args ...string) error {
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

func commandMapBack(cfg *config.Config, cache *pokecache.Cache, args ...string) error {
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

func commandExplore(cfg *config.Config, cache *pokecache.Cache, args ...string) error {
    currAreaURL := *&cfg.LocationAreaURL + args[0]

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
