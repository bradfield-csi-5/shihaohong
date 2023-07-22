package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type server struct {
	location string
	port     string
	time     string
}

/*
go run clockwall/clockwall.go Europe/London=localhost:8010 Asia/Tokyo=localhost:8020 US/Eastern=localhost:8030

stdout:
Europe/London: TBD      |       Asia/Tokyo: TBD |       US/Eastern: TBD |
Europe/London: 16:14:20 |       Asia/Tokyo: 00:14:20    |       US/Eastern: 11:14:20  |
Europe/London: 16:14:21 |       Asia/Tokyo: 00:14:21    |       US/Eastern: 11:14:21  |
Europe/London: 16:14:22 |       Asia/Tokyo: 00:14:22    |       US/Eastern: 11:14:22  |
Europe/London: 16:14:23 |       Asia/Tokyo: 00:14:23    |       US/Eastern: 11:14:23  |
Europe/London: 16:14:24 |       Asia/Tokyo: 00:14:24    |       US/Eastern: 11:14:24  |
*/
func main() {
	args := os.Args[1:]
	servers := make([]server, len(args))

	for i, arg := range args {
		argRes := strings.Split(arg, "=")
		servers[i] = server{
			location: argRes[0],
			port:     argRes[1],
			time:     "TBD",
		}

		go func(s *server) {
			conn, err := net.Dial("tcp", s.port)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()

			sc := bufio.NewScanner(conn)
			for sc.Scan() {
				s.time = sc.Text()
			}
		}(&servers[i])
	}

	for {
		for _, s := range servers {
			fmt.Printf("%s: %s\t|\t", s.location, s.time)
		}
		fmt.Printf("\n")
		time.Sleep(1 * time.Second)
	}
}
