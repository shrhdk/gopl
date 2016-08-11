package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	ports := make([]int, len(os.Args[1:]))
	for i, v := range os.Args[1:] {
		s, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		ports[i] = s
	}

	out := make([]chan string, len(ports))
	for i, port := range ports {
		out[i] = make(chan string)
		fmt.Printf("Connect to %d\n", port)
		go connect(port, out[i])
	}

	for {
		fmt.Println("---------")
		for i, port := range ports {
			fmt.Printf("%d %s", port, <-out[i])
		}
	}
}

func connect(port int, out chan<- string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		out <- string(line)
	}
}
