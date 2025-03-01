package config

type Config struct {
    BaseURL         string
    NextPageURL     *string
    PreviousPageURL *string

    LocationAreaURL string
    PokemonURL      string
}

func NewConfig() *Config {
    return &Config{
        BaseURL: "https://pokeapi.co/api/v2/location-area",
        LocationAreaURL: "https://pokeapi.co/api/v2/location-area/",
        PokemonURL: "https://pokeapi.co/api/v2/pokemon/",
    }
}
