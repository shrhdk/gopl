package main

import (
	"bytes"
	"fmt"
	"log"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
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
	courses, err := topoSort(prereqs)

	if err != nil {
		log.Fatal(err)
	}

	for i, course := range courses {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string, stack []string) error

	visitAll = func(items []string, stack []string) error {
		for _, item := range items {
			for i, stacked := range stack {
				if stacked == item {
					graph := createGraph(append(stack[i:], item))
					return fmt.Errorf("circular dependency is detected.\n\t%s\n", graph)
				}
			}

			if seen[item] {
				continue
			}

			seen[item] = true

			err := visitAll(m[item], append(stack, item))
			if err != nil {
				return err
			}

			order = append(order, item)
		}

		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	err := visitAll(keys, nil)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func createGraph(items []string) string {
	var b bytes.Buffer
	for i, item := range items {
		if i != 0 {
			b.WriteString(" -> ")
		}

		b.WriteString(item)
	}
	return b.String()
}
