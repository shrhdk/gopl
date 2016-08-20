package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	c    chan string
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func after(t time.Duration, f func()) (reset, cancel func()) {
	resetC := make(chan struct{})
	cancelC := make(chan struct{})
	go func() {
		ticker := time.NewTicker(t)
		for {
			select {
			case <-ticker.C:
				f()
				return
			case <-resetC:
				ticker.Stop()
				ticker = time.NewTicker(t)
			case <-cancelC:
				return
			}
		}
	}()

	return func() { resetC <- struct{}{} }, func() { cancelC <- struct{}{} }
}

func participants(clients map[client]bool) string {
	var buf bytes.Buffer
	for client, _ := range clients {
		if buf.String() == "" {
			buf.WriteString("participants: ")
		} else {
			buf.WriteString(", ")
		}
		buf.WriteString(client.name)
	}
	return buf.String()
}

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.c <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			cli.c <- participants(clients)
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.c)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := client{}
	ch.c = make(chan string, 10) // outgoing client messages
	go clientWriter(conn, ch.c)

	input := bufio.NewScanner(conn)

	reset, cancel := after(5*time.Minute, func() {
		fmt.Fprintln(conn, "time out")
		conn.Close()
	})
	defer cancel()

	// determine pariticpant name
	fmt.Fprint(conn, "name: ")
	input.Scan()
	reset()
	ch.name = input.Text()
	ch.c <- "You are " + ch.name

	messages <- ch.name + " has arrived"
	entering <- ch

	for input.Scan() {
		reset()
		messages <- ch.name + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- ch.name + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	host := "localhost"
	port := 8000
	addr := fmt.Sprintf("%s:%d", host, port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("try 'nc %s %d' to connect chat server\n", host, port)

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
