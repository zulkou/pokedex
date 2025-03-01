# Pokedex
## Description
A pokedex simulation CLI app. This project is one of the guided project in [boot.dev](https://www.boot.dev/lessons/5be3e3bd-efb5-4664-a9e9-7111be783271).

## How to Use This Project
Run shell below to download the project:
```bash
$ git clone https://github.com/zulkou/pokedex
$ cd pokedex
```
Make sure to install [go](https://go.dev/). Then you can build the project or run it manually using `go run .`.
How to build the project:
```bash
$ go build
$ ./pokedex
```
## How to Play The Game
After the game is running, you can use `help` command to view the list of available command. There are a few commands:

| Command                   | Description                            |
| ------------------------- | -------------------------------------- |
| `exit`                    | Exit the Pokedex app                   |
| `help`                    | Displays a help message                |
| `map`                     | Displays 20 locations in Pokemon World |
| `mapb`                    | Return to last 20 locations showed     |
| `explore <location-area>` | Explore given location area            |
| `catch <pokemon>`         | Try catch given pokemon                |
| `inspect <pokemon>`       | Inspect caught pokemon                 |
| `pokedex`                 | View pokedex list of caught pokemon    |
