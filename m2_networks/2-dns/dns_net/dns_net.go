package main

import (
	"encoding/binary"
	"strconv"
	// "encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

const (
	SERVER_HOST = "8.8.8.8"
	SERVER_PORT = "53"
	SERVER_TYPE = "udp"
)

const (
	IPv4len                        = 4
	DNSHeaderLength                = 12
	DNSStandardQueryHeaderFlags    = 0x0120 // Wireshark val, TODO: decompose it
	DNSStandardResponseHeaderFlags = 0x8180 // Wireshark val, TODO: decompose it
	DNSMessageMaxLength            = 512    // RFC 1035, 2.3.4
)

type DNSHeader struct {
	ID      uint16
	Flags   uint16
	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

type DNSQuestion struct {
	QName  string
	QType  uint16
	QClass uint16
}

type ResourceRecord struct {
	Name     uint16
	TypeCode uint16
	Class    uint16
	TTL      uint32
	RDLength uint16
	RData    string
}

type DNSMessage struct {
	Header     DNSHeader
	Question   DNSQuestion
	Answers    []ResourceRecord
	Authority  []ResourceRecord
	Additional []ResourceRecord
}

// A domain name represented as a sequence of labels, where
// each label consists of a length octet followed by that
// number of octets.  The domain name terminates with the
// zero length octet for the null label of the root.  Note
// that this field may be an odd number of octets; no
// padding is used.
func (msg *DNSMessage) encodeQName() []byte {
	var data = make([]byte, 0)

	// get qname labels
	labels := strings.Split(msg.Question.QName, ".")

	// encode length
	// encode each character as bytes
	for _, label := range labels {
		data = append(data, byte(len(label)))
		data = append(data, []byte(label)...)
	}

	// encode 0x00 to terminate
	data = append(data, byte(0))
	return data
}

func (msg *DNSMessage) encodeDNSQuery() []byte {
	var data = make([]byte, 0)
	data = binary.BigEndian.AppendUint16(data, msg.Header.ID)
	data = binary.BigEndian.AppendUint16(data, msg.Header.Flags)
	data = binary.BigEndian.AppendUint16(data, msg.Header.QDCount)
	data = binary.BigEndian.AppendUint16(data, msg.Header.ANCount)
	data = binary.BigEndian.AppendUint16(data, msg.Header.NSCount)
	data = binary.BigEndian.AppendUint16(data, msg.Header.ARCount)
	data = append(data, msg.encodeQName()...)
	data = binary.BigEndian.AppendUint16(data, msg.Question.QType)
	data = binary.BigEndian.AppendUint16(data, msg.Question.QClass)
	return data
}

func decodeDNSMessage(b []byte) DNSMessage {
	header := DNSHeader{
		ID:      binary.BigEndian.Uint16(b[0:2]),
		Flags:   binary.BigEndian.Uint16(b[2:4]),
		QDCount: binary.BigEndian.Uint16(b[4:6]),
		ANCount: binary.BigEndian.Uint16(b[6:8]),
		NSCount: binary.BigEndian.Uint16(b[8:10]),
		ARCount: binary.BigEndian.Uint16(b[10:12]),
	}

	qName, qNameLen := decodeQName(b[12:])
	currentByteOffset := DNSHeaderLength + qNameLen
	question := DNSQuestion{
		QName:  qName,
		QType:  binary.BigEndian.Uint16(b[currentByteOffset : currentByteOffset+2]),
		QClass: binary.BigEndian.Uint16(b[currentByteOffset+2 : currentByteOffset+4]),
	}
	currentByteOffset += 4
	answers := make([]ResourceRecord, header.ANCount)
	for i := uint16(0); i < header.ANCount; i++ {
		RDLen := binary.BigEndian.Uint16(b[currentByteOffset+10 : currentByteOffset+12])
		typeCode := binary.BigEndian.Uint16(b[currentByteOffset+2 : currentByteOffset+4])

		var rData string
		if typeCode == 1 {
			rData = strconv.Itoa(int(b[currentByteOffset+12])) + "." + strconv.Itoa(int(b[currentByteOffset+13])) + "." + strconv.Itoa(int(b[currentByteOffset+14])) + "." + strconv.Itoa(int(b[currentByteOffset+15]))
		} else if typeCode == 5 {
			rData = decodeCName(b[currentByteOffset+12:], int(RDLen))
		} else {
			rData = string(b[currentByteOffset+12 : currentByteOffset+12+int(RDLen)])
		}

		answers[i] = ResourceRecord{
			Name:     binary.BigEndian.Uint16(b[currentByteOffset : currentByteOffset+2]),
			TypeCode: typeCode,
			Class:    binary.BigEndian.Uint16(b[currentByteOffset+4 : currentByteOffset+6]),
			TTL:      binary.BigEndian.Uint32(b[currentByteOffset+6 : currentByteOffset+10]),
			RDLength: RDLen,
			RData:    rData,
		}
		currentByteOffset += 12 + int(RDLen)
	}

	return DNSMessage{
		Header:   header,
		Question: question,
		Answers:  answers,
	}
}

// Assume b[0] is the start of QName
func decodeQName(b []byte) (qname string, len int) {
	qname = ""
	len = 0
	labelLen := b[len]
	for {
		len++
		for j := uint8(0); j < labelLen; j++ {
			qname += string(b[len])
			len++
		}

		labelLen = b[len]
		if labelLen == 0 {
			len++
			break
		} else {
			qname += "."
		}
	}
	return
}

// Assume b[0] is the start of CName
// TODO: implement offset checking (c0)
func decodeCName(b []byte, len int) (cname string) {
	cname = ""
	i := 0
	for i < len {
		labelLen := int(b[i])
		// TODO: Making bad assumption ".com" here
		if labelLen == 0xc0 {
			cname += "com"
			break
		}

		i++
		end := i + labelLen
		for j := i; j < end; j++ {
			cname += string(b[j])
		}
		cname += "."
		i += labelLen
	}
	fmt.Println(cname)

	return
}

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
