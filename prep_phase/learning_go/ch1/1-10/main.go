// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	url_pkg "net/url"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	secs := time.Since(start).Seconds()

	if err != nil {
		ch <- fmt.Sprintf("while reading all %s, %v", url, err)
	}
	resp.Body.Close() // don't leak resources

	parsedUrl, err := url_pkg.Parse(url)
	urlString := strings.TrimPrefix(parsedUrl.Hostname(), "www.")
	if err != nil {
		ch <- fmt.Sprintf("while parsing url %s, %v", url, err)
	}

	err = os.WriteFile("output_"+urlString, body, 0644)
	if err != nil {
		ch <- fmt.Sprintf("while writing file %s, %v", url, err)
	}

	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, len(body), url)
}
