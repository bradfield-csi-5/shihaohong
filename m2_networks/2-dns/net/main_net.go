package main

import (
	// "encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	url := os.Args[1]

	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

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

	_, err = connection.Write(query.encodeDNSQuery())
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, DNSMessageMaxLength)
	mLen, err := connection.Read(buffer)
	if err != nil {
		panic(err)
	}

	response := decodeDNSMessage(buffer[:mLen])
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
