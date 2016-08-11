package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func echo(c net.Conn, shout string, delay time.Duration) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c *net.TCPConn) {
	defer func() {
		c.Close()
		log.Println("close")
	}()

	in := scanner(c)
	tiemr := time.NewTimer(10 * time.Second)

	for {
		select {
		case s := <-in:
			tiemr.Stop()
			tiemr = time.NewTimer(10 * time.Second)
			wg.Add(1)
			go echo(c, s, 1*time.Second)
		case <-tiemr.C:
			tiemr.Stop()
			fmt.Fprintln(c, "time out")
			return
		}
	}
}

func scanner(c *net.TCPConn) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		in := bufio.NewScanner(c)
		for in.Scan() {
			ch <- in.Text()
		}
	}()
	return ch
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
