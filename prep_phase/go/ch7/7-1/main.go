package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Counter struct {
	count     int
	splitFunc bufio.SplitFunc
}

func (c *Counter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(strings.NewReader(string(p)))
	s.Split(c.splitFunc)

	for s.Scan() {
		c.count++
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	return c.count, nil
}

func main() {
	// test WordCounter
	str := "lorem ipsum text blah blah blah   hahaha"
	var wc = Counter{
		splitFunc: bufio.ScanWords,
	}

	fmt.Fprintf(&wc, "input string: %s", str)
	fmt.Printf("word count: %d\n", wc.count)
	fmt.Fprintf(&wc, "more strings")
	fmt.Printf("word count: %d\n", wc.count)
	res, err := wc.Write([]byte("interesting content here"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("word count: %d\n", res)

	// test LineCounter
	str = "lorem ipsum text \nblah blah blah   hahaha"
	var lc = Counter{
		splitFunc: bufio.ScanLines,
	}

	fmt.Fprintf(&lc, "input string: %s", str)
	fmt.Printf("line count: %d\n", lc.count)
	fmt.Fprintf(&lc, "more strings")
	fmt.Printf("line count: %d\n", lc.count)
	res, err = lc.Write([]byte("interesting \ncontent \nhere"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("line count: %d\n", res)
}
