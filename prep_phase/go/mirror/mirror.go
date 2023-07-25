package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// where the mirror files will be put into
const basename = "contents"

func extractLinks(resp *http.Response) ([]string, error) {
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", resp.Request.URL, err)
	}

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
func processUrl(url string) {
	// get file
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Printf("getting %s: %s", url, resp.Status)
		return
	}

	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		resp.Body.Close()
		log.Printf("content %s is not text/html: %s", url, resp.Header.Get("Content-Type"))
		return
	}

	// save into local disk
	filepath := basename + resp.Request.URL.Path
	createDirectory(filepath)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		log.Println(err)
		return
	}
	createFile(filepath, "index.html", data)

	links, err := extractLinks(resp)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(links)
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

	// init mirror directory
	createDirectory(basename)

	// process first file
	processUrl(args[0])
}
