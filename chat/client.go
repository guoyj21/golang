package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var running bool

func receive(conn *net.Conn) {
	for running {
		fmt.Println(readFromConn(conn))
		fmt.Print("You> ")
	}
}

func readFromConn(conn *net.Conn) string {
	buf := make([]byte, 2048)
	_, err := (*conn).Read(buf)
	if err != nil {
		(*conn).Close()
		running = false
		return "Reading from connection error"
	}

	str := string(buf)
	fmt.Println()
	return str
}

func send(conn *net.Conn) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("You> ")
		input, _ := reader.ReadBytes('\n')
		if bytes.Equal(input, []byte("quit\n")) {
			(*conn).Write([]byte("/quit"))
			running = false
			break
		}

		(*conn).Write(input[0 : len(input)-1])
	}
}

func main() {
	running = true

	serverAddr := "127.0.0.1:8888"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatal("Connect to server error")
	}

	fmt.Println("Please input your name")
	reader := bufio.NewReader(os.Stdin)

	name, _ := reader.ReadBytes('\n')

	conn.Write(name[0 : len(name)-1])

	go receive(&conn)

	go send(&conn)

	for running {
		time.Sleep(time.Second)
	}

}
