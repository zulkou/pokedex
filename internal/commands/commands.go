package commands

import (
	"fmt"
	"os"

	"github.com/zulkou/pokedex/internal/api"
	"github.com/zulkou/pokedex/internal/config"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(cfg *config.Config) error
}

var Commands map[string]CliCommand

func CommandExit(cfg *config.Config) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func CommandHelp(cfg *config.Config) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
    for _, v := range Commands {
        fmt.Printf("%v: %v\n", v.Name, v.Description)
    }
    return nil
}

func CommandMap(cfg *config.Config) error {
    urlToUse := cfg.BaseURL
    if cfg.NextPageURL != nil {
        urlToUse = *cfg.NextPageURL
    }

    data, err := api.FetchLocation(urlToUse)
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

func CommandMapBack(cfg *config.Config) error {
    if cfg.PreviousPageURL == nil {
        fmt.Println("You're on the first page, there's no going back!")
        return nil
    }

    data, err := api.FetchLocation(*cfg.PreviousPageURL)
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

