// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
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

// !+
func handleConn(c net.Conn, cnt chan<- struct{}) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		cnt <- struct{}{}
		go echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

//!-

type Err struct {
	Error error
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	cnt := make(chan struct{})
	go func(l net.Listener, cnt chan<- struct{}) {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Fatal(err)
			}
			go handleConn(conn, cnt)
		}
	}(l, cnt)

	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		log.Printf("time remaining: %d", countdown)
		select {
		case <-tick:
		case <-cnt:
			countdown = 10
		}
	}
	log.Printf("server shutting down\n")
}
