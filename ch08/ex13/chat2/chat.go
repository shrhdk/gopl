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

func timeoutCloser(conn net.Conn, timeout time.Duration) (reset, cancel func()) {
	resetC := make(chan struct{})
	cancelC := make(chan struct{})
	go func() {
		ticker := time.NewTicker(timeout)
		for {
			select {
			case <-ticker.C:
				fmt.Fprintln(conn, "time out")
				conn.Close()
			case <-resetC:
				ticker.Stop()
				ticker = time.NewTicker(timeout)
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
	ch.c = make(chan string) // outgoing client messages
	go clientWriter(conn, ch.c)

	who := conn.RemoteAddr().String()
	ch.c <- "You are " + who
	ch.name = who
	messages <- who + " has arrived"
	entering <- ch

	reset, cancel := timeoutCloser(conn, 5*time.Second)
	defer cancel()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		reset()
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

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
