package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("https://example.org")
	if err != nil {
		log.Fatal(err)
	}
	rel, err := u.Parse("https://google.com/foo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rel)
	_, err = u.Parse(":foo")
	if _, ok := err.(*url.Error); !ok {
		log.Fatal(err)
	}
}
