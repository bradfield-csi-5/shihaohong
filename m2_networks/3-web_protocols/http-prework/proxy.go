package main

import (
	"fmt"
	"net"
	"strings"

	"golang.org/x/sys/unix"
)

const (
	SERVER_PORT      = 3001
	SERVER_BACKLOG   = 10
	DST_FORWARD_PORT = 9000
)

var SERVER_HOST = [4]byte{127, 0, 0, 1}
var cache = NewCache()

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseRequestPath(req string) string {
	s := strings.Split(req, "\r\n")
	return s[0]
}

func main() {
	// socket to the internet
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, unix.IPPROTO_TCP)
	check(err)
	defer unix.Close(fd)

	internetSocket := &unix.SockaddrInet4{
		Port: SERVER_PORT,
		Addr: SERVER_HOST,
	}
	err = unix.Bind(fd, internetSocket)
	check(err)

	fmt.Println("listening on ", SERVER_HOST, ":", SERVER_PORT)
	err = unix.Listen(fd, SERVER_BACKLOG)
	check(err)

	for {
		nfd, _, err := unix.Accept(fd)
		check(err)
		go handleConnection(nfd)
	}
}

func handleConnection(fd int) {
	defer unix.Close(fd)
	buf := make([]byte, 1024)
	n, _, err := unix.Recvfrom(fd, buf, 0)
	check(err)
	if n <= 0 {
		return
	}

	fmt.Println("Request from client:")
	fmt.Println(string(buf[:n]))

	fmt.Println("Request path:")
	fmt.Println(parseRequestPath(string(buf[:n])))

	path := parseRequestPath(string(buf[:n]))

	res := cache.Get(path)
	if len(res) > 0 {
		fmt.Println("Cached:")
		fmt.Println(string(res))
		err = unix.Send(fd, res, 0)
		check(err)
		return
	}

	// socket to the other port
	// TODO: figure out how to implement this using unix.Socket()
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   SERVER_HOST[:],
		Port: DST_FORWARD_PORT,
	})
	check(err)

	fmt.Printf("connection to forward port %s success\n", conn.LocalAddr().String())
	defer conn.Close()

	_, err = conn.Write(buf[:n])
	check(err)

	fwdBuf := make([]byte, 1024)
	fwdN, err := conn.Read(fwdBuf)
	check(err)

	fmt.Println("Response from server:")
	fmt.Println(string(fwdBuf[:fwdN]))
	cache.Put(path, fwdBuf[:fwdN])

	err = unix.Send(fd, fwdBuf, 0)
	check(err)
}
