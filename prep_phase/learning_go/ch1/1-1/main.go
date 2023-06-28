package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	Echo()
	Echo2()
	JoinEcho()
}

func Echo() {
	for i, val := range os.Args {
		fmt.Printf("%d:\t%s\n", i, val)
	}
}

func JoinEcho() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
