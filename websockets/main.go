package main

import (
	"bufio"
	"net"
	"encoding/json"
	"fmt"
	"net/http"
)

type Packet struct {
	...
}

// Creating chanenel
type Channel struct {
	conn net.Conn		// Base websocket connection
	send chan Packet	// Outgoing packets queue
}


func main() {
	fmt.Print
	json.Unmarshal
	http.CanonicalHeaderKey
}

func NewChannel(conn net.Conn) *Channel {
	c := &Channel{
		conn : conn, 
		send: make(chan Packet ,N)
	}

	go c.reader()		// Launching read go routines that runs indefinitely
	go c.writer()		// Launching write routing that runs indefinitely

	return c
}

func (c *Channel) reader() {
	// Making a buffered read call to reduce syscalls.
	buf := bufio.NewReader(c.conn)

	// Unlimited loop set up since we expect data to come in
	for {
		pkt, _ := readPacket(buf)
		c.handle(pkt)
	}
}
func (c *Channel) writer() {
	buf := bufio.NewWriter(c.conn)

	for pkt := range c.send {
		_ := writePacket(buf, pkt)
		buf.Flush()
	}

}
