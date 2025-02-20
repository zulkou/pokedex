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
    commands.InitializeCommand()

    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        input := scanner.Text()

        if command, exists := commands.Commands[input]; exists {
            command.Callback(conf, cache)
            continue;
        } else {
            fmt.Println("Unknown command")
        }

        // clean_input := cleanInput(input)
        // fmt.Printf("Your command was: %v\n", clean_input[0])
    }
}
