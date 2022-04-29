package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTPClient is a http client.
type HTTPClient interface {
	Get(path string) ([]byte, error)
}

// Client is the client for the dictionary API.
type Client struct {
	Key string
}

const (
	BaseURL = "https://www.dictionaryapi.com/api/v3/references"
)

// NewClient returns a new Client.
func NewClient(key string) *Client {
	return &Client{Key: key}
}

// Get returns the response for the given word.
func (c *Client) Get(word string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/collegiate/json/%s?key=%s", BaseURL, word, c.Key))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read body, err: %w", err)
	}

	return b, nil
}
