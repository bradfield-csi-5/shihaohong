package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter struct {
	count  int64
	writer io.Writer
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	if err != nil {
		return 0, err
	}

	c.count += int64(n)
	return n, nil
}

// Returns a new io.Writer wrapping the original, with an int64 variable that contains
// num of bytes written to new writer
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	bc := ByteCounter{
		count:  0,
		writer: w,
	}
	return &bc, &bc.count
}

// code from https://golang.cafe/blog/golang-writer-example.html
func main() {
	w, err := os.OpenFile("123.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// wrap os.OpenFile with ByteCounter
	cw, cnt := CountingWriter(w)

	// test that write does as intended
	n, err := cw.Write([]byte("writing some data into a file\n"))
	fmt.Printf("wrote %d bytes\n", n)
	if err != nil {
		panic(err)
	}

	// test a second write call
	n, err = cw.Write([]byte("more bytes into the file\n"))
	fmt.Printf("wrote %d bytes\n", n)
	if err != nil {
		panic(err)
	}

	// validate that the counter is incrementing accordingly
	fmt.Printf("total %d bytes written\n", *cnt)
}
