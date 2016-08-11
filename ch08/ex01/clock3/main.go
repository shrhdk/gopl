package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		help()
	}

	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		help()
	}

	loc, err := time.LoadLocation(os.Args[2])
	if err != nil {
		help()
	}

	listen(port, loc)
}

func help() {
	fmt.Printf("[   Usage   ] %s <port> <location>\n", os.Args[0])
	fmt.Printf("[ Example 1 ] %s 8000 Asia/Tokyo\n", os.Args[0])
	fmt.Printf("[ Example 2 ] %s 8000 America/Chicago\n", os.Args[0])
	os.Exit(1)
}

func listen(port int, loc *time.Location) {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn, loc)
	}
}

func handleConn(c net.Conn, loc *time.Location) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format(time.RFC1123)+"\n")
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
