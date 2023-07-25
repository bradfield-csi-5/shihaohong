// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

func crawl(it item) []item {
	list, err := links.Extract(it.url)
	if err != nil {
		log.Print(err)
	}
	items := []item{}
	for _, url := range list {
		items = append(items, item{url, it.depth + 1})
	}
	return items
}

type item struct {
	url   string
	depth int
}

// !+
func main() {
	depthFlag := flag.Int("depth", 3, "depth of links to crawl for")
	flag.Parse()

	if *depthFlag <= 0 {
		panic("-depth flag must be a non-zero positive int")
	}

	worklist := make(chan []item)  // lists of URLs, may have duplicates
	unseenLinks := make(chan item) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func(d int) {
		args := flag.Args()
		items := []item{}
		for _, arg := range args {
			items = append(items, item{arg, d})
		}
		worklist <- items
	}(1)

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, item := range list {
			fmt.Printf("url: %s\tdepth:%d\n", item.url, item.depth)
			if !seen[item.url] {
				seen[item.url] = true
				if item.depth < *depthFlag {
					unseenLinks <- item
				}
			}
		}
	}
}

//!-
