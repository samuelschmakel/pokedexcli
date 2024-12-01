package pokeapi

import (
	"io"
	"net/http"
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

func GetLocationAreas(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
