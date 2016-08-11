package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileBytesChans := make([]chan int64, len(roots))
	for i, root := range roots {
		fileBytesChans[i] = make(chan int64)
		walkDir2(root, fileBytesChans[i])
	}

	dus := make([]diskusage, len(roots))

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}

	alive := len(roots)
loop:
	for {
		for i := 0; i < len(roots); i++ {
			select {
			case size, ok := <-fileBytesChans[i]:
				if !ok {
					alive--
					break // fileSizes was closed
				}
				dus[i].nfiles++
				dus[i].nbytes += size
			case <-tick:
				printAllDiskUsage(roots, dus)
			}

			if alive == 0 {
				break loop
			}
		}
	}
	printAllDiskUsage(roots, dus)

}

type diskusage struct {
	nfiles int64
	nbytes int64
}

func walkDir2(dir string, fileSizes chan<- int64) {
	var n sync.WaitGroup

	n.Add(1)
	go walkDir(dir, &n, fileSizes)

	go func() {
		n.Wait()
		close(fileSizes)
	}()
}

func printAllDiskUsage(dirs []string, dus []diskusage) {
	for i := 0; i < len(dirs); i++ {
		printDiskUsage(dirs[i], dus[i].nfiles, dus[i].nbytes)
	}
}

func printDiskUsage(dir string, nfiles, nbytes int64) {
	fmt.Printf("%s %d files  %.1f GB\n", dir, nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{} // acquire token
	defer func() {
		<-sema
	}() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
