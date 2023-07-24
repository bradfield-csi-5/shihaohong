// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 227.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

// modified to close only the write half of the connection so that when
// the program will continue to print the final echoes from reverb1 server
// after stdinput has been closed
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()

	// copies from stdin and sends to server
	mustCopy(conn, os.Stdin)

	// this works with ctrl + d (send EOF signal) because
	// it'll close the write side of the TCP connection
	// on hitting EOF in the stdin
	conn.(*net.TCPConn).CloseWrite()
	// conn.Close()

	<-done // wait for background goroutine to finish
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
