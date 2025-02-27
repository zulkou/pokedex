package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/zulkou/pokedex/internal/commands"
	"github.com/zulkou/pokedex/internal/config"
	"github.com/zulkou/pokedex/internal/pokecache"
)

func cleanInput(text string) []string {
    lowered := strings.ToLower(text)
    trimmed := strings.Trim(lowered, " ")
    split := strings.Fields(trimmed)
    return split
}

func main() {
    conf := config.NewConfig()
    cache := pokecache.NewCache(30 * time.Second)
    defer cache.Close()
    commands.InitializeCommand()

    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        input := strings.Split(scanner.Text(), " ")

        main := input[0]
        args := input[1:]

        if command, exists := commands.Commands[main]; exists {
            err := command.Callback(conf, cache, args...)
            if err != nil {
                fmt.Println(err)
            }
            continue;
        } else {
            fmt.Println("Unknown command")
        }

        // clean_input := cleanInput(input)
        // fmt.Printf("Your command was: %v\n", clean_input[0])
    }
}
