package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	messages        = make(chan string)
)

var (
	host = flag.String("h", "localhost", "host")
	port = flag.Int("p", 3090, "port")
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	message := make(chan string)
	go MessageWrite(conn, message)

	clientName := conn.RemoteAddr().String()

	message <- fmt.Sprintf("Welcome to the server, your name %s", clientName)

	messages <- fmt.Sprintf("New client is here, name %s", clientName)
	incomingClients <- message

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s:%d", clientName, inputMessage.Text())
	}

	leavingClients <- message
	messages <- fmt.Sprintf("%s said bye", clientName)
}

func MessageWrite(conn net.Conn, message <-chan string) {
	for message := range messages {
		fmt.Fprintln(conn, message)
	}
}
