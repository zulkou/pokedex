package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/zulkou/pokedex/internal/pokecache"
)

type LocationArea struct {
    Results []struct {
        Name string `json:"name"`
    } `json:"results"`
    Next        *string `json:"next"`
    Previous    *string `json:"previous"`
}

func FetchLocation(url string, cache *pokecache.Cache) (*LocationArea, error) {
    if data, found := cache.Get(url); found {
        var location LocationArea
        err := json.Unmarshal(data, &location)
        return &location, err
    }

    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error making API request: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err!= nil {
        return nil, fmt.Errorf("error reading response body: %w", err)
    }

    cache.Add(url, body)

    var result LocationArea
    err = json.Unmarshal(body, &result)
    if err != nil {
        return nil, fmt.Errorf("error parsing json: %w", err)
    }

    return &result, nil
}
