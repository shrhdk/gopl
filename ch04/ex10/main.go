package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shrhdk/gopl/ch04/ex10/github"
)

func main() {
	searchAndPrintIssues("In less than a month", combine([]string{"created:>" + yyyymmdd(-31)}, os.Args[1:]))
	searchAndPrintIssues("In less than a year old", combine([]string{"created:>" + yyyymmdd(-365)}, os.Args[1:]))
	searchAndPrintIssues("In more than a year old", combine([]string{"created:<" + yyyymmdd(-365)}, os.Args[1:]))
}

func searchAndPrintIssues(title string, terms []string) {
	result, err := github.SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\n=== %s ===\n\n", title)
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func yyyymmdd(diffDay int) string {
	const format = "2006-01-02"
	t := time.Now()
	t = t.AddDate(0, 0, diffDay)
	return t.Format(format)
}

func combine(s1, s2 []string) []string {
	c := make([]string, len(s1)+len(s2))
	copy(c[:len(s1)], s1)
	copy(c[len(s1):], s2)
	return c
}
