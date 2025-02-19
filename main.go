package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zulkou/pokedex/internal/commands"
	"github.com/zulkou/pokedex/internal/config"
)

func InitializeCommand() {
    commands.Commands = map[string]commands.CliCommand{
        "exit": {
            Name:        "exit",
            Description: "Exit the Pokedex",
            Callback:    commands.CommandExit,
        },
        "help": {
            Name: "help",
            Description: "Displays a help message",
            Callback: commands.CommandHelp,
        },
        "map": {
            Name: "map",
            Description: "Displays 20 locations in Pokemon World",
            Callback: commands.CommandMap,
        },
        "mapb": {
            Name: "mapb",
            Description: "Return to last 20 locations showed",
            Callback: commands.CommandMapBack,
        },
    }
}

func cleanInput(text string) []string {
    lowered := strings.ToLower(text)
    trimmed := strings.Trim(lowered, " ")
    split := strings.Fields(trimmed)
    return split
}

func main() {
    config := config.NewConfig()
    InitializeCommand()

    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        input := scanner.Text()

        if command, exists := commands.Commands[input]; exists {
            command.Callback(config)
            continue;
        } else {
            fmt.Println("Unknown command")
        }

        // clean_input := cleanInput(input)
        // fmt.Printf("Your command was: %v\n", clean_input[0])
    }
}
