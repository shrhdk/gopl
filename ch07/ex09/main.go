package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

// Data

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func newTracks() []*Track {
	return []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
}

// Sorter

type byKey struct {
	t       []*Track
	sortKey string
}

func (x byKey) Len() int { return len(x.t) }

func (x byKey) Less(i, j int) bool {
	a, b := x.t[i], x.t[j]
	switch x.sortKey {
	case "Title":
		if a.Title != b.Title {
			return a.Title < b.Title
		}
	case "Artist":
		if a.Artist != b.Artist {
			return a.Artist < b.Artist
		}
	case "Album":
		if a.Album != b.Album {
			return a.Album < b.Album
		}
	case "Year":
		if a.Year != b.Year {
			return a.Year < b.Year
		}
	case "Length":
		if a.Length != b.Length {
			return a.Length < b.Length
		}
	}
	return false
}

func (x byKey) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

// Web Server

var tracksList = template.Must(template.New("tracksList").Parse(`<!DOCTYPE html>
<html>
<head>
	<title>Track List</title>
</head>
<body>
	<style>table {border-collapse: collapse;} th,td {padding: 0.5em; border: black 1px solid;}</style>
	<h1>Track List</h1>
	<table>
		<tr>
			<th><a href="?sort_key=Title">Title</a></th>
			<th><a href="?sort_key=Artist">Artist</a></th>
			<th><a href="?sort_key=Album">Album</a></th>
			<th><a href="?sort_key=Year">Year</a></th>
			<th><a href="?sort_key=Length">Length</a></th>
		</tr>
		{{range .}}
		<tr>
		  <td>{{.Title}}</td>
		  <td>{{.Artist}}</td>
		  <td>{{.Album}}</td>
		  <td>{{.Year}}</td>
		  <td>{{.Length}}</td>
		</tr>
		{{end}}
	</table>
	</body>
</html>`))

func main() {
	sockAddress := "localhost:8000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sortKey := r.FormValue("sort_key")
		tracks := newTracks()
		sort.Sort(byKey{tracks, sortKey})
		if err := tracksList.Execute(w, tracks); err != nil {
			log.Fatal(err)
		}
	})
	fmt.Println("Serving on " + sockAddress)
	log.Fatal(http.ListenAndServe(sockAddress, nil))
	return
}
