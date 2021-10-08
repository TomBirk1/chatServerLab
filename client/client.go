package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn *net.Conn) {
	for {
		reader := bufio.NewReader(*conn)
		msg, _ := reader.ReadString('\n')
		fmt.Printf(msg)
	}
	//TODO In a continuous loop, read a message from the server and display it.
}

func write(conn *net.Conn) {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter message: ")
		text, _ := stdin.ReadString('\n')

		fmt.Fprintln(*conn, text)
		read(conn)
	}
	//TODO Continually get input from the user and send messages to the server.
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	conn, _ := net.Dial("tcp", "127.0.0.1:8030")
	read(&conn)
	write(&conn)
	//TODO Try to connect to the server
	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
}
