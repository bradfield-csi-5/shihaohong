package main

import (
	"fmt"
	"strconv"

	"github.com/popcount/utils"
)

func main() {
	binaryStr := "10101010"
	num, err := strconv.ParseUint(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error parsing binary string:", err)
		return
	}

	uintVal := uint64(num)
	fmt.Println(utils.PopCount(uintVal))
	fmt.Println(utils.PopCountLoop(uintVal))
	fmt.Println(utils.PopCountShift(uintVal))
	fmt.Println(utils.PopCountClear(uintVal))
}
