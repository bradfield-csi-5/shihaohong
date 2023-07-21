package main

import (
	"bufio"
	"fmt"
	"strings"
)

type LineCounter int
type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(strings.NewReader(string(p)))
	count, err := countUsingSplit(s, bufio.ScanWords)
	if err != nil {
		return 0, err
	}

	*wc += WordCounter(count)
	return len(p), nil
}

func (lc *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(strings.NewReader(string(p)))
	count, err := countUsingSplit(s, bufio.ScanLines)
	if err != nil {
		return 0, err
	}

	*lc += LineCounter(count)
	return len(p), nil
}

func countUsingSplit(s *bufio.Scanner, sf bufio.SplitFunc) (int, error) {
	s.Split(sf)

	count := 0
	for s.Scan() {
		count++
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

func main() {
	// test WordCounter
	str := "lorem ipsum text blah blah blah   hahaha"
	var wc WordCounter

	fmt.Println("===WordCounter===")
	fmt.Fprintf(&wc, "input string: %s", str)
	fmt.Printf("word count: %d\n", wc)
	fmt.Fprintf(&wc, "more strings")
	fmt.Printf("word count: %d\n", wc)
	res, err := wc.Write([]byte("interesting content here"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("num bytes read: %d\n", res)
	fmt.Printf("word count: %d\n", wc)

	// test LineCounter
	str = "lorem ipsum text \nblah blah blah   hahaha"
	var lc LineCounter

	fmt.Println("===LineCounter===")
	fmt.Fprintf(&lc, "input string: %s", str)
	fmt.Printf("line count: %d\n", lc)
	fmt.Fprintf(&lc, "more strings")
	fmt.Printf("line count: %d\n", lc)
	res, err = lc.Write([]byte("interesting \ncontent \nhere"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("num bytes read: %d\n", res)
	fmt.Printf("line count: %d\n", lc)
}
