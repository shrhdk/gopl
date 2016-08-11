package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for item, _ := range sliceToSet(items) {
			if seen[item] {
				continue
			}

			seen[item] = true
			visitAll(m[item])
			order = append(order, item)
		}
	}

	keys := getKeys(m)

	visitAll(keys)
	return order
}

func getKeys(m map[string][]string) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func sliceToSet(s []string) map[string]struct{} {
	m := make(map[string]struct{})

	for _, key := range s {
		m[key] = struct{}{}
	}

	return m
}
