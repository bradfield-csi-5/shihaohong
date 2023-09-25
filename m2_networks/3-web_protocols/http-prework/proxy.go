package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

const (
	SERVER_PORT    = 3000
	SERVER_BACKLOG = 10
)

var SERVER_HOST = [4]byte{127, 0, 0, 1}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, unix.IPPROTO_TCP)
	check(err)
	svrSockInet4 := &unix.SockaddrInet4{
		Port: SERVER_PORT,
		Addr: SERVER_HOST,
	}
	err = unix.Bind(fd, svrSockInet4)
	check(err)
	fmt.Println("listening on ", SERVER_HOST, ":", SERVER_PORT)
	err = unix.Listen(fd, SERVER_BACKLOG)
	check(err)

	buf := make([]byte, 512)
	for {
		nfd, _, err := unix.Accept(fd)
		check(err)

		n, _, err := unix.Recvfrom(nfd, buf, 0)
		check(err)

		err = unix.Send(nfd, buf[:n], 0)
		check(err)

		unix.Close(nfd)
	}
}
