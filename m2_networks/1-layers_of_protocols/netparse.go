package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"os"
	"sort"
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

// TODO: check if URG set and add urgent pointer
type TCPHeader struct {
	SourcePort            uint16
	DestinationPort       uint16
	SequenceNumber        uint32
	AcknowledgementNumber uint32
	DataOffset            uint8 // in bytes
	Reserved              uint8 // only 4 bits
	Flags                 uint8
	WindowSize            uint16
	Checksum              uint16
}

// Assume URG not set (manually checked)
// TODO: handle urgent pointer dynamically
func NewTCPHeader(bs []byte) (TCPHeader, error) {
	return TCPHeader{
		SourcePort:            binary.BigEndian.Uint16(bs[:2]),
		DestinationPort:       binary.BigEndian.Uint16(bs[2:4]),
		SequenceNumber:        binary.BigEndian.Uint32(bs[4:8]),
		AcknowledgementNumber: binary.BigEndian.Uint32(bs[8:12]),
		DataOffset:            (bs[12] & 0xf0) >> 2, // >> 4 to get the value, then << 2 to get in 4-byte words
		Reserved:              bs[12] & 0xf,
		Flags:                 bs[13],
		WindowSize:            binary.BigEndian.Uint16(bs[14:16]),
		Checksum:              binary.BigEndian.Uint16(bs[16:18]),
	}, nil
}

type HTTPPacket struct {
	Sequence uint32
	Data     []byte
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

	fmt.Println("processing packets")
	packetCount := 0

	httpPackets := make([]HTTPPacket, 0, fiLen)
	seqCheck := make(map[uint32]struct{})
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

		bp := fp + EthernetHeaderLength
		eh, err := NewEthernetHeader(fi[fp:bp])
		if err != nil {
			panic(err)
		}

		// Assumes IPv4 header.
		// TODO: Can modify to check for both and
		// branch it, but that's too much work for now
		ih, err := NewIPv4Header(fi[bp:])
		if err != nil {
			panic(err)
		}

		bp += int(ih.IHL)

		tcph, err := NewTCPHeader(fi[bp:])
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
		fmt.Println("TCP HEADER DETAILS:")
		fmt.Printf("SourcePort: %d\n", tcph.SourcePort)
		fmt.Printf("DestinationPort: %d\n", tcph.DestinationPort)
		fmt.Printf("TransportHeaderLength: %d\n", tcph.DataOffset)
		fmt.Printf("SequenceNumber: %d\n", tcph.SequenceNumber)

		seq := tcph.SequenceNumber
		if tcph.SourcePort == 80 {
			if _, ok := seqCheck[seq]; ok {
				// skip is somehow duplicate packet sequence found
				fp += int(packetHeader.Length)
				continue
			}
			seqCheck[seq] = struct{}{}
			httpPackets = append(httpPackets, HTTPPacket{
				Sequence: seq,
				Data:     fi[bp+int(tcph.DataOffset):],
			})
		}

		fp += int(packetHeader.Length)
		packetCount++
	}

	sort.Slice(httpPackets, func(i, j int) bool {
		return httpPackets[i].Sequence < httpPackets[j].Sequence
	})

	// open output file
	fo, err := os.Create("output.jpeg")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	for _, httpPacket := range httpPackets {
		if _, err := fo.Write(httpPacket.Data); err != nil {
			panic(err)
		}
	}
}
