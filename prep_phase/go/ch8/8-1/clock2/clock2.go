package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn, lf string) {
	defer c.Close()
	for {
		loc, _ := time.LoadLocation(lf)
		now := time.Now().In(loc)
		_, err := io.WriteString(c, now.Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

/*
go run clock2.go -tz=Europe/London -port 8010 & go run clock2.go -tz=Asia/Tokyo -port 8020 & go run clock2.go -tz=US/Eastern -port 8030 &
then, run/build clockwall
*/
func main() {
	portFlag := flag.Int("port", 8000, "the port to start this clock server on")
	tzFlag := flag.String("tz", "Local", "the location to return the time for. ie. \"Asia/Tokyo\", \"US/Eastern\"")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *portFlag))

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, *tzFlag) // handle connections concurrently
	}
}
