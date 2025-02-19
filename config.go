package main

type Config struct {
    BaseURL         string
    NextPageURL     *string
    PreviousPageURL *string
}

func NewConfig() *Config {
    return &Config{
        BaseURL: "https://pokeapi.co/api/v2/location-area",
    }
}
