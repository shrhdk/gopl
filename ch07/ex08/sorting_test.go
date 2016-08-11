package ex08

import (
	"fmt"
	"sort"
	"testing"
)

func newTracks() []*Track {
	return []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
}

func Test(t *testing.T) {
	var track []*Track

	track = newTracks()
	s, sc := CountingSort(ByKeys(track, "Title", "Year"))
	sort.Sort(s)
	fmt.Printf("By Simple Sort (%s)\n", sc)
	printTracks(track)

	fmt.Println("")

	track = newTracks()
	count := MultiStableSort(track, "Title", "Year")
	fmt.Printf("By Multiple Stable Sort (%s)\n", count)
	printTracks(track)
}
