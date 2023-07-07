// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/sha256popcount/utils"
)

func main() {
	// all matching
	c1 := sha256.Sum256([]byte("a"))
	c2 := sha256.Sum256([]byte("a"))
	fmt.Printf("c1: %x\nc2: %x\n", c1, c2)
	res := utils.CountMatchingBits(c1, c2)
	fmt.Println("num matching bits in sha256 vals: ", res)

	c1 = sha256.Sum256([]byte("qwerty"))
	c2 = sha256.Sum256([]byte("qwerty"))
	fmt.Printf("\nc1: %x\nc2: %x\n", c1, c2)
	res = utils.CountMatchingBits(c1, c2)
	fmt.Println("num matching bits in sha256 vals: ", res)

	// different vals
	c1 = sha256.Sum256([]byte("a"))
	c2 = sha256.Sum256([]byte("b"))
	fmt.Printf("\nc1: %x\nc2: %x\n", c1, c2)
	res = utils.CountMatchingBits(c1, c2)
	fmt.Println("num matching bits in sha256 vals: ", res)

	c1 = sha256.Sum256([]byte("fus"))
	c2 = sha256.Sum256([]byte("ro dah!"))
	fmt.Printf("\nc1: %x\nc2: %x\n", c1, c2)
	res = utils.CountMatchingBits(c1, c2)
	fmt.Println("num matching bits in sha256 vals: ", res)
}
