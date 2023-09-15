package main

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const GlobalHeaderLength = 24

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

const filename = "net.cap"

func main() {
	// TODO: stream file rather than loading entire file
	// to memory
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
		packetHeader := fi[fp : fp+16]
		fp += 16
		// fmt.Println("packet header val: 0x" + hex.EncodeToString(packetHeader))

		packetLen := packetHeader[8:12]
		_, packetLenDec := parsePacketLen(packetLen)

		untruncatedPacketLen := (packetHeader[12:])
		_, untruncatedPacketLenDec := parsePacketLen(untruncatedPacketLen)

		if packetLenDec != untruncatedPacketLenDec {
			fmt.Printf("packetLen: %v\n", packetLenDec)
			fmt.Printf("untruncatedPacketLen: %v\n", untruncatedPacketLenDec)
			panic("something's wrong packetLen != untruncatedPacketLen")
		}

		packets = append(packets, fi[fp:fp+packetLenDec]...)
		fp += packetLenDec
	}

	fmt.Println("packets val: 0x" + hex.EncodeToString(packets[:80]))
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

func parsePacketLen(s []byte) (hexStr string, decInt int) {
	reverseByteOrdering(s) // fix byte ordering
	hexStr = hex.EncodeToString(s)
	dec, err := strconv.ParseInt(hexStr, 16, 32)
	if err != nil {
		panic(err)
	}
	decInt = int(dec)
	return
}

func reverseByteOrdering(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
