package main

import (
	"bufio"
	"fmt"
	"os"
)

// go run main.go < input.txt
func main() {
	res := wordfreq()
	fmt.Println(res)

	for c, n := range res {
		fmt.Printf("%q:\t%d\n", c, n)
	}
}

func wordfreq() map[string]int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	res := make(map[string]int)
	for scanner.Scan() {
		res[scanner.Text()]++
	}

	return res
}
