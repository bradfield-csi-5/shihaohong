package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// where the mirror files will be put into
const basename = "contents"

// counting semaphore for how many parallel operations can be open
var tokens = make(chan struct{}, 20)

func extractLinks(resp *http.Response, data []byte) ([]string, error) {
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", resp.Request.URL, err)
	}

	// TODO: URLs need to be altered to point to the mirrored page
	// use a set to avoid duplication
	links := map[string]struct{}{}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}

				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}

				link.Fragment = ""
				if link.Host == resp.Request.Host {
					links[link.String()] = struct{}{}
				}
			}
		}
	}
	forEachNode(doc, visitNode)

	linkSlice := []string{}
	for k := range links {
		linkSlice = append(linkSlice, k)
	}
	return linkSlice, nil
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre)
	}
}

// Gets the document, puts it into a root directory appending the relative path
// of the document.
func processUrl(url string) []string {
	tokens <- struct{}{}
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Printf("getting %s: %s", url, resp.Status)
		return nil
	}

	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		resp.Body.Close()
		log.Printf("content %s is not text/html: %s", url, resp.Header.Get("Content-Type"))
		return nil
	}

	// save into local disk
	filepath := basename + resp.Request.URL.Path
	createDirectory(filepath)

	data, err := io.ReadAll(resp.Body)
	resp.Body = io.NopCloser(bytes.NewBuffer(data))
	if err != nil {
		resp.Body.Close()
		log.Println(err)
		return nil
	}
	createFile(filepath, "index.html", data)

	links, err := extractLinks(resp, data)
	<-tokens
	if err != nil {
		log.Println(err)
		return nil
	}

	return links
}

func createDirectory(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func createFile(path string, filename string, data []byte) {
	err := os.WriteFile(path+"/"+filename, data, 0666)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		log.Println("Please pass in a single URL to process. For example, https://go.dev")
	}

	createDirectory(basename)
	worklist := make(chan []string)
	var n int
	n++
	go func(link string) {
		links := processUrl(link)
		worklist <- links
	}(args[0])

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				fmt.Printf("url: %s\n", link)
				seen[link] = true
				n++
				go func(link string) {
					links := processUrl(link)
					worklist <- links
				}(link)
			}
		}
	}
}
