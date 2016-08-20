package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"gopl.shiro.be/ch04/ex14/github"
)

func printUsageAndExit() {
	fmt.Fprintf(os.Stderr, "%s owner/reponame username accessToken", os.Args[0])
	os.Exit(1)
}

var tmpl = template.Must(template.New("repoinfo").Parse(`<!DOCTYPE html>
<html>
<head>
	<title>{{.Owner}}/{{.Name}}</title>
</head>
<body>
	<h1>github.com/{{.Owner}}/{{.Name}}</h1>
	<h2>Issues</h2>
	<ul>
	{{range .Issues}}
		<li><a href="{{.URL}}">{{.Title}}</a></li>
	{{end}}
	</ul>
	<h2>Milestones</h2>
	<ul>
	{{range .Milestones}}
		<li><a href="{{.URL}}">{{.Title}}</a></li>
	{{end}}
	</ul>
	<h2>Collaborators</h2>
	<ul>
	{{range .Collaborators}}
		<li><a href="{{.URL}}">{{.Login}}</a></li>
	{{end}}
	</ul>
</body>
</html>`))

func startServer(repo *github.Repository) {
	http.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, repo)
	}))

	sockAddress := "localhost:8000"
	fmt.Println("Serving on " + sockAddress)
	log.Fatal(http.ListenAndServe(sockAddress, nil))
}

func main() {
	if len(os.Args) != 4 {
		printUsageAndExit()
	}

	reponame := strings.Split(os.Args[1], "/")
	if len(reponame) != 2 {
		printUsageAndExit()
	}

	owner := reponame[0]
	name := reponame[1]
	user := os.Args[2]
	accessToken := os.Args[3]

	repo, err := github.FetchRepository(owner, name, user, accessToken)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(repo)

	startServer(repo)
}
