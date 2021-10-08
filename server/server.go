package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strconv"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	for {
		conn, _ := ln.Accept()
		conns <- conn

	}
	// TODO: all
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {

	for {
		reader := bufio.NewReader(client)
		msg, _ := reader.ReadString('\n')
		message := strconv.Itoa(clientid) + " : " + msg
		newMessage := Message{sender: clientid, message: message}
		msgs <- newMessage
	}
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	//TODO Create a Listener for TCP connections on the port given above.
	ln, _ := net.Listen("tcp", *portPtr)
	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)
	i := 0

	//Start accepting connections
	go acceptConns(ln, conns)

	for {
		select {
		case conn := <-conns:
			clients[i] = conn
			go handleClient(conn, i, msgs)
			i++

			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients channel
			// - start to asynchronously handle messages from this client
		case msg := <-msgs:
			for j := 0; j < i; j++ {
				if j != msg.sender {
					fmt.Fprint(clients[j], msg.message)
				}
			}
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender
		}
	}
}
