package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{
		Title: "Casablanca",
		Year: 1942,
		Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
	},
	{
		Title: "Cool Hand Luke",
		Year: 1967,
		Color: true,
		Actors: []string{"Paul Newman"},
	},
	{
		Title: "Bullitt",
		Year: 1968,
		Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"},
	},
}

func main() {
	// [{Casablanca 1942 false [Humphrey Bogart Ingrid Bergman]} {Cool Hand Luke 1967 true [Paul Newman]} {Bullitt 1968 true [Steve McQueen Jacqueline Bisset]}]
	fmt.Println(movies)

	data, err := json.Marshal(movies)
	if err != nil {
		fmt.Println("wocao")
	}
	fmt.Printf("%s\n", data)
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}