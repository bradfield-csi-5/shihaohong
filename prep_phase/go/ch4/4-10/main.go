// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"gopl.io/ch4/github"
)

// !+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	items := result.Items

	sort.Slice(items, func(i, j int) bool {
		return items[i].CreatedAt.After(items[j].CreatedAt)
	})

	now := time.Now()
	// durations used to compare. making assumptions about number of days a month
	monthDuration := time.Hour * 24 * 30
	yearDuration := monthDuration * 12

	var i int
	fmt.Println("===within the month===")
	for i < len(items) {
		item := items[i]
		if now.Sub(item.CreatedAt) > monthDuration {
			i++
			break
		}

		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
		i++
	}

	fmt.Println("")
	fmt.Println("===within the year===")
	for i < len(items) {
		item := items[i]
		if now.Sub(item.CreatedAt) > yearDuration {
			i++
			break
		}

		fmt.Printf("%s #%-5d %9.9s %.55s \n",
			item.CreatedAt.Format("2006-01-02"), item.Number, item.User.Login, item.Title)
		i++
	}

	fmt.Println("")
	fmt.Println("===older than a year===")
	for i < len(items) {
		item := items[i]
		fmt.Printf("%s #%-5d %9.9s %.55s \n",
			item.CreatedAt.Format("2006-01-02"), item.Number, item.User.Login, item.Title)
		i++
	}
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
