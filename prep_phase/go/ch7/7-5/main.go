package main

import (
	"fmt"
	"io"
	"os"
)

type LimitRead struct {
	r           io.Reader
	limit       int64
	alreadyRead int64
}

func (lr *LimitRead) Read(p []byte) (n int, err error) {
	len := int64(len(p))
	remainingBytes := lr.limit - lr.alreadyRead
	if len < remainingBytes {
		n, err := lr.r.Read(p)
		if err != nil {
			return 0, err
		}
		return n, nil
	} else {
		n, err := lr.r.Read(p[:remainingBytes])
		if err != nil {
			return 0, err
		}

		if len == remainingBytes {
			return n, nil
		}

		return n, io.EOF
	}
}

func LimitReader(r io.Reader, n int64) io.Reader {
	lr := LimitRead{
		r:           r,
		limit:       n,
		alreadyRead: 0,
	}

	return &lr
}

func main() {
	r, err := os.Open("123.txt")
	if err != nil {
		panic(err)
	}

	// test happy path
	cr := LimitReader(r, 100)
	b1 := make([]byte, 80)

	n1, err := cr.Read(b1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	r.Close()

	// test exact bytes
	r, err = os.Open("123.txt")
	if err != nil {
		panic(err)
	}
	cr2 := LimitReader(r, 10)
	b2 := make([]byte, 10)

	n2, err := cr2.Read(b2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes: %s\n", n2, string(b2[:n2]))
	r.Close()

	r, err = os.Open("123.txt")
	if err != nil {
		panic(err)
	}
	cr3 := LimitReader(r, 10)
	b3 := make([]byte, 11)
	n3, err := cr3.Read(b3)
	if err == io.EOF {
		fmt.Println(`reached EOF, but expected`)
	} else if err != nil {
		// unexpected error
		panic(err)
	}
	// bytes still read
	fmt.Printf("%d bytes: %s\n", n3, string(b3[:n3]))
}
