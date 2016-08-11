package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	comics, err := fetchAll([]int{1, 2, 3, 4, 5}...)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	fmt.Println("ready\n")

	for _, comic := range comics {
		fmt.Printf("hint: %s\n", comic.Title)
	}
	fmt.Println("")

	for {
		var title string
		fmt.Scan(&title)
		num, comic := search(comics, title)
		if comic == nil {
			fmt.Printf("not found\n")
			continue
		}
		fmt.Printf("URL: %s\n", comicURL(num))
		fmt.Printf("Title: %s\n", comic.Title)
	}
}

func search(comics map[int]Comic, title string) (int, *Comic) {
	for num, comic := range comics {
		if strings.Contains(comic.Title, title) {
			return num, &comic
		}
	}
	return 0, nil
}

func fetchAll(nums ...int) (map[int]Comic, error) {
	comics := make(map[int]Comic)
	for _, num := range nums {
		comic, err := fetch(num)
		if err != nil {
			return nil, err
		}
		comics[num] = *comic
	}
	return comics, nil
}

func fetch(num int) (*Comic, error) {
	resp, err := http.Get(comicURL(num))

	if err != nil {
		resp.Body.Close()
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, err
	}

	var result Comic
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

func comicURL(num int) string {
	return fmt.Sprintf("http://xkcd.com/%d/info.0.json", num)
}

type Comic struct {
	Title      string
	Transcript string
}
