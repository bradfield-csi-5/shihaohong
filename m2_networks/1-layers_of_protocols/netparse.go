package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

const filename = "net.cap"

func main() {
	// TODO: stream file rather than loading entire file
	// to memory
	fi, err := os.ReadFile(filename)
	fiLen := len(fi)
	if err != nil {
		fmt.Println(err)
	}

	fp := 0
	// capture file header
	globalHeader := fi[:24]
	fmt.Println("global header val: 0x" + hex.EncodeToString(globalHeader))
	fp += 24

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
