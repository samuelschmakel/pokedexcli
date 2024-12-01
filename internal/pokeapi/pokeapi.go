package pokeapi

import (
	"io"
	"net/http"
	"time"

	pokecache "github.com/samuelschmakel/pokedexcli/internal/pokecache"
)

type Config struct {
	Next     string
	Previous string
}

// Internal
func initializeConfig() *Config {
	initialURL := "https://pokeapi.co/api/v2/location-area/"
	return &Config{Next: initialURL, Previous: ""}
}

var InitialConfig = initializeConfig()

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var myCache = pokecache.NewCache(5 * time.Second)

func GetLocationAreas(url string) ([]byte, error) {
	// Check the cache, return the data if it's there
	if data, found := myCache.Get(url); found {
		return data, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// io.ReadAll() reads all the data from an io.Reader into a byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Store in the cache
	myCache.Add(url, body)

	return body, err
}
