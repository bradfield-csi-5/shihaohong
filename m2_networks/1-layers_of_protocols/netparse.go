package main

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	GlobalHeaderLength = 24
	PacketHeaderLength = 16
)

var ErrInvalidHeaderLength = errors.New("invalid header length")

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
	// capture packet header
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
		fp += int(packetHeader.Length)
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
