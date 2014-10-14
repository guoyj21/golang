package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
)

type ClientChat struct {
	name        string
	conn        *net.Conn
	OUT         chan string
	broadcastCh chan string
	quit        chan bool
	clientChain []ClientChat
}

func (c *ClientChat) Read(buf []byte) bool {
	nr, err := (*c.conn).Read(buf)

	if err != nil {
		(*c.conn).Close()
		return false
	}

	fmt.Println("Read(): ", nr, " bytes")
	return true
}

func (c *ClientChat) Close() {
	c.quit <- true
	(*c.conn).Close()

}

func (c *ClientChat) Equal(c1 *ClientChat) bool {
	if bytes.Equal([]byte(c.name), []byte(c1.name)) {
		if c.conn == c1.conn {
			return true
		}
	}
	return false
}

func (c *ClientChat) deleteFromList() {
	//a = append(a[:i], a[i+1:]...)
	for i, _ := range c.clientChain {
		temp := c.clientChain[i]
		if c.Equal(&temp) {
			c.clientChain = append(c.clientChain[:i], c.clientChain[i+1:]...)
		}
	}
}

func broadcaster(broadcastCh <-chan string, clientList *[]ClientChat) {
	for {
		input := <-broadcastCh

		for _, value := range *clientList {
			//notice all client
			ClientChat(value).OUT <- input
		}
	}
}

func receive(client *ClientChat) {
	buf := make([]byte, 2048)

	for client.Read(buf) {
		if bytes.Equal(buf, []byte("/quit")) {
			client.Close()
			break
		}

		send := client.name + "> " + string(buf)
		client.broadcastCh <- send

		for i := 0; i < 2048; i++ {
			buf[i] = 0x00
		}
	}

	client.broadcastCh <- client.name + " has left chat"
}

func send(client *ClientChat) {
	for {
		select {
		case buf := <-client.OUT:
			(*client.conn).Write([]byte(buf))
		case <-client.quit:
			(*client.conn).Close()
			break
		}
	}
	fmt.Println("send(): stop for: ", client.name)
}

func clientHandler(conn *net.Conn, broadcastCh chan string, clientChain *[]ClientChat) {
	buf := make([]byte, 1024)
	(*conn).Read(buf)
	name := string(buf)
	newClient := ClientChat{name, conn, make(chan string), broadcastCh, make(chan bool), *clientChain}

	go send(&newClient)
	go receive(&newClient)

	*clientChain = append(*clientChain, newClient)
	broadcastCh <- name + " has join the chat"

}

func main() {
	clientList := []ClientChat{}

	broadcastCh := make(chan string)

	go broadcaster(broadcastCh, &clientList)

	netListen, err := net.Listen("tcp", "127.0.0.1:8888")

	if err != nil {
		log.Fatal("Listen socket error")
		os.Exit(-1)
	}

	for {
		log.Println("Wait for Client...")
		conn, err := netListen.Accept()
		if err != nil {
			log.Println("Accept client error")
			continue
		}

		go clientHandler(&conn, broadcastCh, &clientList)
	}
}
