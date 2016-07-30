package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	OMDBURL = "https://www.omdbapi.com/?r=json"
)

type Movie struct {
	Title  string
	Poster string
}

func SearchMovies(terms []string) (*Movie, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(OMDBURL + "&t=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result Movie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
