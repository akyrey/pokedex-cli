package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/akyrey/pokedex-cli/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient *http.Client
}

func NewHttpClient(interval, timeout time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) ListLocations(pageUrl *string) (*Pagination[NameUrlStruct], error) {
	url := baseUrl + "/location-area/"
	if pageUrl != nil {
		url = *pageUrl
	}

	if cached, ok := c.cache.Get(url); ok {
		var p Pagination[NameUrlStruct]
		err := json.Unmarshal(cached, &p)
		if err != nil {
			return nil, err
		}

		return &p, nil
	}

	r, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed retrieving location areas: %s", r.Status)
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var p Pagination[NameUrlStruct]
	err = json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, data)
	return &p, nil
}

func (c *Client) ExploreLocation(area string) (*ExploreResponse, error) {
	url := baseUrl + "/location-area/" + area

	if cached, ok := c.cache.Get(url); ok {
		var explore ExploreResponse
		err := json.Unmarshal(cached, &explore)
		if err != nil {
			return nil, err
		}

		return &explore, nil
	}

	r, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed retrieving pokémons from location: %s", r.Status)
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var explore ExploreResponse
	err = json.Unmarshal(data, &explore)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, data)
	return &explore, nil
}

func (c *Client) GetPokemon(name string) (*Pokemon, error) {
	url := baseUrl + "/pokemon/" + name

	if cached, ok := c.cache.Get(url); ok {
		var pokemon Pokemon
		err := json.Unmarshal(cached, &pokemon)
		if err != nil {
			return nil, err
		}

		return &pokemon, nil
	}

	r, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed retrieving pokémon: %s", r.Status)
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, data)
	return &pokemon, nil
}
