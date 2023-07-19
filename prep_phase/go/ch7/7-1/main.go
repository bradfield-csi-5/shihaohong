package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(strings.NewReader(string(p)))
	s.Split(bufio.ScanWords)

	for s.Scan() {
		*wc = *wc + WordCounter(1)
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	return int(*wc), nil
}

type LineCounter int

func (lc *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(strings.NewReader(string(p)))
	s.Split(bufio.ScanLines)

	for s.Scan() {
		*lc = *lc + LineCounter(1)
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	return int(*lc), nil
}

func main() {
	// test WordCounter
	str := "lorem ipsum text blah blah blah   hahaha"
	var wc WordCounter

	fmt.Fprintf(&wc, "input string: %s", str)
	fmt.Printf("word count: %d\n", wc)
	fmt.Fprintf(&wc, "more strings")
	fmt.Printf("word count: %d\n", wc)
	res, err := wc.Write([]byte("interesting content here"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("word count: %d\n", res)

	// test LineCounter
	str = "lorem ipsum text \nblah blah blah   hahaha"
	var lc LineCounter

	fmt.Fprintf(&lc, "input string: %s", str)
	fmt.Printf("line count: %d\n", lc)
	fmt.Fprintf(&lc, "more strings")
	fmt.Printf("line count: %d\n", lc)
	res, err = lc.Write([]byte("interesting \ncontent \nhere"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("line count: %d\n", res)
}
