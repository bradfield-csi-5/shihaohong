package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"os"
)

const (
	GlobalHeaderLength   = 24
	PacketHeaderLength   = 16
	EthernetHeaderLength = 14
)

var ErrInvalidHeaderLength = errors.New("invalid header length")
var ErrInvalidIPHeader = errors.New("expected ipv4 segments")

type GlobalHeader struct {
	MagicNumber      []byte
	MajorVersion     uint16
	MinorVersion     uint16
	TimezoneAccuracy uint32
	TimezoneOffset   uint32
	SnapshotLength   uint32
	LinkType         uint32
}

func NewGlobalHeader(bs []byte) (GlobalHeader, error) {
	if len(bs) != GlobalHeaderLength {
		return GlobalHeader{}, ErrInvalidHeaderLength
	}
	return GlobalHeader{
		MagicNumber:      bs[:4],
		MajorVersion:     binary.LittleEndian.Uint16(bs[4:6]),
		MinorVersion:     binary.LittleEndian.Uint16(bs[6:8]),
		TimezoneAccuracy: binary.LittleEndian.Uint32(bs[8:12]),
		TimezoneOffset:   binary.LittleEndian.Uint32(bs[12:16]),
		SnapshotLength:   binary.LittleEndian.Uint32(bs[16:20]),
		LinkType:         binary.LittleEndian.Uint32(bs[20:24]),
	}, nil
}

type PacketHeader struct {
	TimestampSec      uint32
	TimestampMsNs     uint32
	Length            uint32
	UntruncatedLength uint32
}

func NewPacketHeader(bs []byte) (PacketHeader, error) {
	if len(bs) != PacketHeaderLength {
		return PacketHeader{}, ErrInvalidHeaderLength
	}
	return PacketHeader{
		TimestampSec:      binary.LittleEndian.Uint32(bs[0:4]),
		TimestampMsNs:     binary.LittleEndian.Uint32(bs[4:8]),
		Length:            binary.LittleEndian.Uint32(bs[8:12]),
		UntruncatedLength: binary.LittleEndian.Uint32(bs[12:16]),
	}, nil
}

type EthernetHeader struct {
	MACDest   net.HardwareAddr
	MACSrc    net.HardwareAddr
	EtherType uint16
}

func NewEthernetHeader(bs []byte) (EthernetHeader, error) {
	if len(bs) != EthernetHeaderLength {
		return EthernetHeader{}, ErrInvalidHeaderLength
	}
	return EthernetHeader{
		MACDest:   bs[0:6],
		MACSrc:    bs[6:12],
		EtherType: binary.BigEndian.Uint16(bs[12:14]),
	}, nil
}

// See http://www.tcpipguide.com/free/t_IPDatagramGeneralFormat.htm
type IPv4Header struct {
	Version            uint8
	IHL                uint8 // length in bytes
	TypeOfService      byte
	TotalLength        uint16
	Identification     uint16
	FlagsAndFragments  []byte
	TTL                uint8
	Protocol           uint8
	HeaderChecksum     uint16
	SourceAddress      IPAddress
	DestinationAddress IPAddress
}

type IPAddress struct {
	b []byte
}

func (ip *IPAddress) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", int(ip.b[0]), int(ip.b[1]), int(ip.b[2]), int(ip.b[3]))
}

func NewIPv4Header(bs []byte) (IPv4Header, error) {
	ipVersion := uint8(bs[0] >> 4)
	if ipVersion != 4 {
		return IPv4Header{}, ErrInvalidIPHeader
	}

	ihl := uint8((bs[0] & 0xF) << 2) // shift <<2 bc it's in multiples of 4-byte words
	return IPv4Header{
		Version:            ipVersion,
		IHL:                ihl,
		TypeOfService:      bs[1],
		TotalLength:        binary.BigEndian.Uint16(bs[2:4]),
		Identification:     binary.BigEndian.Uint16(bs[4:6]),
		FlagsAndFragments:  bs[7:8],
		TTL:                uint8(bs[8]),
		Protocol:           uint8(bs[9]),
		HeaderChecksum:     binary.BigEndian.Uint16(bs[10:12]),
		SourceAddress:      IPAddress{bs[12:16]},
		DestinationAddress: IPAddress{bs[16:20]},
	}, nil
}

func main() {
	// TODO: stream file rather than loading entire file
	// to memory
	const filename = "net.cap"
	fi, err := os.ReadFile(filename)
	fiLen := len(fi)
	if err != nil {
		panic(err)
	}

	fp := 0
	// capture file header
	_, err = NewGlobalHeader(fi[:GlobalHeaderLength])
	if err != nil {
		panic(err)
	}
	// fmt.Println("0x" + hex.EncodeToString(globalHeader.MagicNumber)) // 0xd4c3b2a1
	fp += GlobalHeaderLength

	packets := make([]byte, 0, fiLen)
	fmt.Println("processing packets")
	packetCount := 0
	for fp < fiLen {
		packetHeader, err := NewPacketHeader(fi[fp : fp+PacketHeaderLength])
		if err != nil {
			panic(err)
		}
		fp += PacketHeaderLength
		if packetHeader.Length != packetHeader.UntruncatedLength {
			fmt.Printf("packetLen: %v\n", packetHeader.Length)
			fmt.Printf("untruncatedPacketLen: %v\n", packetHeader.UntruncatedLength)
			panic("something's wrong packetLen != untruncatedPacketLen")
		}

		packets = append(packets, fi[fp:fp+int(packetHeader.Length)]...)

		ip := fp + EthernetHeaderLength
		eh, err := NewEthernetHeader(fi[fp:ip])
		if err != nil {
			panic(err)
		}

		// Assumes IPv4 header.
		// TODO: Can modify to check for both and
		// branch it, but that's too much work for now
		ih, err := NewIPv4Header(fi[ip:])
		if err != nil {
			panic(err)
		}

		fmt.Printf("==== packet %d ====\n", packetCount)
		fmt.Println("ETHERNET HEADER DETAILS:")
		fmt.Printf("MAC address src: %s\n", eh.MACSrc.String())
		fmt.Printf("MAC address dest: %s\n", eh.MACDest.String())
		fmt.Println("IP HEADER DETAILS:")
		fmt.Printf("IP address src: %s\n", ih.SourceAddress.String())
		fmt.Printf("IP address dest: %s\n", ih.DestinationAddress.String())
		fmt.Printf("datagram length: %d\n", ih.TotalLength)
		fmt.Printf("IP header length: %d\n", ih.IHL) // constant 20, no optional headers
		fmt.Printf("protocol: %d\n", ih.Protocol)    // UDP protocol

		fp += int(packetHeader.Length)
		packetCount++
	}

	dumpEthernetPacketData(packets)
}

func dumpEthernetPacketData(arr []byte) {
	f, err := os.Create("ethernet.bin")
	if err != nil {
		panic("Couldn't open file")
	}
	defer f.Close()

	err = binary.Write(f, binary.LittleEndian, arr)
	if err != nil {
		panic("Write failed")
	}
}
