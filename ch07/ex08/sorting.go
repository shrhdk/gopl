package ex08

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

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

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

// ByKeys

type byKeys struct {
	t        []*Track
	sortKeys []string
}

func (x byKeys) Len() int { return len(x.t) }

func (x byKeys) Less(i, j int) bool {
	a, b := x.t[i], x.t[j]
	for _, sortKey := range x.sortKeys {
		switch sortKey {
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
	}
	return false
}

func (x byKeys) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func ByKeys(t []*Track, sortKeys ...string) sort.Interface {
	return byKeys{t, sortKeys}
}

// CountingSort

type SortCount struct {
	nLen, nLess, nSwap int64
}

func (sc SortCount) String() string {
	return fmt.Sprintf("nLen: %d, nLess: %d, nSwap: %d", sc.nLen, sc.nLess, sc.nSwap)
}

func (sc *SortCount) add(a SortCount) {
	sc.nLen += a.nLen
	sc.nLess += a.nLess
	sc.nSwap += a.nSwap
}

type countingSort struct {
	base sort.Interface
	SortCount
}

func (x *countingSort) Len() int {
	x.nLen++
	return x.base.Len()
}

func (x *countingSort) Less(i, j int) bool {
	x.nLess++
	return x.base.Less(i, j)
}

func (x *countingSort) Swap(i, j int) {
	x.nSwap++
	x.base.Swap(i, j)
}

func CountingSort(baseSort sort.Interface) (sort.Interface, *SortCount) {
	cs := countingSort{baseSort, SortCount{0, 0, 0}}
	return &cs, &(cs.SortCount)
}

// MultiStableSort

func MultiStableSort(t []*Track, sortKeys ...string) SortCount {
	var sortCount SortCount
	for i := len(sortKeys) - 1; i >= 0; i-- {
		s, sc := CountingSort(ByKeys(t, sortKeys[i]))
		sort.Stable(s)
		sortCount.add(*sc)
	}
	return sortCount
}
