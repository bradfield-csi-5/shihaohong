package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func loadCSV(filename string, maxLoad int) []Tuple {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(f)
	res := []Tuple{}

	// read table fields/attributes
	colNames, _ := reader.Read()
	count := 0
	for {
		if count == maxLoad {
			break
		}
		count++
		cols, err := reader.Read()
		values := []Value{}
		for i, c := range cols {
			values = append(values, Value{
				Name:        colNames[i],
				StringValue: c,
			})
		}

		res = append(res, Tuple{Values: values})
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}

	return res
}

func main() {
	log.Println("loading csv")
	tuples := loadCSV("movies.csv", 10)
	log.Println("csv load complete")

	scanOp := NewScanOperator(tuples)
	limitOp := NewLimitOperator(5, scanOp)

	// All tuples should be returned.
	for range tuples {
		fmt.Println(limitOp.Next())
		fmt.Println(limitOp.Execute())
	}
}
