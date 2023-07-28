package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)

	ticker := time.NewTicker(1 * time.Second)
	ec := make(chan struct{})
	go func(c net.Conn, input *bufio.Scanner, e chan<- struct{}) {
		for input.Scan() {
			e <- struct{}{}
			go echo(c, input.Text(), 1*time.Second)
		}
	}(c, input, ec)

	for countdown := 10; countdown > 0; countdown-- {
		log.Printf("connection %v countdown: %d\n", c.RemoteAddr(), countdown)
		select {
		case <-ticker.C:
		case <-ec:
			countdown = 10
		}
	}

	// NOTE: ignoring potential errors from input.Err()
	log.Printf("closing connection %v", c.RemoteAddr())
	ticker.Stop()
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
