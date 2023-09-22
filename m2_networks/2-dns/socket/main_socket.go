package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"math/rand"
	"os"
	"time"
)

var SERVER_HOST = [4]byte{8, 8, 8, 8}

func main() {
	url := os.Args[1]

	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, unix.IPPROTO_UDP)
	if err != nil {
		panic(err)
	}

	// hard-code query for now
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	query := DNSMessage{
		Header: DNSHeader{
			ID:      uint16(r1.Uint32() % 0xFFFF),
			Flags:   DNSStandardQueryHeaderFlags,
			QDCount: 1,
			NSCount: 0,
			ANCount: 0,
			ARCount: 0,
		},
		Question: DNSQuestion{
			QName:  url,
			QType:  1, // assume Type A only for now
			QClass: 1, // https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml
		},
	}

	svrSockInet4 := &unix.SockaddrInet4{
		Port: SERVER_PORT,
		Addr: SERVER_HOST,
	}

	// send DNS message to server using fd provided by socket syscall
	// through the sendto syscall
	err = unix.Sendto(fd, query.encodeDNSQuery(), 0, svrSockInet4)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, DNSMessageMaxLength)
	// use recvfrom syscall to receive the response from the server
	n, recvSockAddr, err := unix.Recvfrom(fd, buffer, 0)
	fmt.Printf("sender address %v\n", recvSockAddr)
	if err != nil {
		panic(err)
	}

	// no bytes received, peer has performed an orderly shudown or -1
	// if waiting on response and its nonblocking in which errno is set to EAGAIN
	if n <= 0 {
		return
	}

	response := decodeDNSMessage(buffer[:n])
	fmt.Println(response)
	// Pretty print in JSON format.
	// RData is not printed correctly as the answer results need to be properly processed.
	// queryJSON, err := json.MarshalIndent(query, "", "  ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Request:\n%s\n", string(queryJSON))
	// responseJSON, err := json.MarshalIndent(response, "", "  ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Response:\n%s\n", string(responseJSON))
}
