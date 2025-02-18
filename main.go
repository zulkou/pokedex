package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandMap map[string]cliCommand

func initializeCommand() {
    commandMap = map[string]cliCommand{
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "help": {
            name: "help",
            description: "Displays a help message",
            callback: commandHelp,
        },
    }
}

func commandExit() error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp() error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
    for _, v := range commandMap {
        fmt.Printf("%v: %v\n", v.name, v.description)
    }
    return nil
}

func cleanInput(text string) []string {
    lowered := strings.ToLower(text)
    trimmed := strings.Trim(lowered, " ")
    split := strings.Fields(trimmed)
    return split
}

func main() {
    /*
    result := cleanInput("Hello, world!  ")
    for i := range result {
        fmt.Println(result[i])
    }
    */
    initializeCommand()

    scanner := bufio.NewScanner(os.Stdin)

    for {
        // fmt.Print("Pokedex > ")
        scanner.Scan()
        input := scanner.Text()

        if command, exists := commandMap[input]; exists {
            command.callback()
            continue;
        } else {
            fmt.Println("Unknown command")
        }

        // clean_input := cleanInput(input)
        // fmt.Printf("Your command was: %v\n", clean_input[0])
    }
}
