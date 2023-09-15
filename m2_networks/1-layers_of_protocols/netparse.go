package main

import (
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

	packetCount := 0
	// capture packet header
	for fp < fiLen {
		packetHeader := fi[fp : fp+16]
		fp += 16
		fmt.Println("packet header val: 0x" + hex.EncodeToString(packetHeader))

		// parse len
		packetLen := packetHeader[8:12]
		reverseByteOrdering(packetLen) // fix byte ordering
		packetLenHex := hex.EncodeToString(packetLen)
		packetLenDec, err := strconv.ParseInt(packetLenHex, 16, 32)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Printf("packetLenHex: %s\n", packetLenHex)
		// fmt.Printf("packetLenDec: %v\n", packetLenDec)

		untruncatedPacketLen := (packetHeader[12:])
		reverseByteOrdering(untruncatedPacketLen) // fix byte ordering
		untruncatedPacketLenHex := hex.EncodeToString(untruncatedPacketLen)
		untruncatedPacketLenDec, err := strconv.ParseInt(untruncatedPacketLenHex, 16, 32)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Printf("untruncatedPacketLenHex: %s\n", untruncatedPacketLenHex)
		// fmt.Printf("untruncatedPacketLenDec: %v\n", untruncatedPacketLenDec)

		if packetLenDec != untruncatedPacketLenDec {
			fmt.Println("something's wrong packetLen != untruncatedPacketLen")
			fmt.Printf("packetLen: %v\n", packetLenDec)
			fmt.Printf("untruncatedPacketLen: %v\n", untruncatedPacketLenDec)
		}

		// TODO: process packet payload

		packetCount++
		fp += int(packetLenDec)
	}

	// fmt.Printf("packet count: %v\n", packetCount) // to confirm packet count
}

func reverseByteOrdering(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
