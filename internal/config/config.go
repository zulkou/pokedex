package config

type Config struct {
    BaseURL         string
    NextPageURL     *string
    PreviousPageURL *string

    LocationAreaURL string
}

func NewConfig() *Config {
    return &Config{
        BaseURL: "https://pokeapi.co/api/v2/location-area",
        LocationAreaURL: "https://pokeapi.co/api/v2/location-area/",
    }
}
