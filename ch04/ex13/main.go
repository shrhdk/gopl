package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const baseURL = "http://www.omdbapi.com/"

func main() {
	title := strings.Join(os.Args[1:], " ")

	movie, err := GetMovie(title)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if movie.Poster == "" {
		fmt.Fprintf(os.Stderr, "No Poster Data\n")
		os.Exit(1)
	}

	GetPoster(os.Stdout, movie.Poster)
}

func GetPoster(o io.Writer, posterURL string) error {
	resp, err := http.Get(posterURL)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("search query failed %s", resp.Status)
	}

	_, err = io.Copy(o, resp.Body)
	if err != nil {
		resp.Body.Close()
		return err
	}

	resp.Body.Close()
	return nil
}

func GetMovie(title string) (*Movie, error) {
	t := url.QueryEscape(title)

	resp, err := http.Get(baseURL + "?t=" + t)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed %s", resp.Status)
	}

	var result Movie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

type Movie struct {
	Title  string
	Poster string
}
