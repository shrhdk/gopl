package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.shiro.be/ch04/ex11/github"
)

func main() {
	sub := os.Args[1]

	switch sub {
	case "create":
		title := os.Args[2]
		body := os.Args[3]
		if err := github.CreateIssue(title, body); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
		break
	case "read":
		num, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}

		issue, err := github.ReadIssue(num)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}

		fmt.Printf("%s\n%s\n", issue.Title, issue.Body)
		break
	case "update":
		num, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
		title := os.Args[3]
		body := os.Args[4]
		if err := github.UpdateIssue(num, title, body); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
	case "close":
		num, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
		if err := github.CloseIssue(num); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
	}
}

func showUsage() {
	fmt.Print(fmt.Sprintf("Usage\n\t%s create|read|update|close [issue-number]\n", os.Args[0]))
	os.Exit(1)
}
