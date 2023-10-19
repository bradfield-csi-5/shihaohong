package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

const src string = `package foo

import (
	"fmt"
	"time"
)

func baz() {
	fmt.Println("Hello, world!")
}

type A int

const b = "testing"

func bar() {
	fmt.Println(time.Now())
}`

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByName []dst.FuncDecl

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name.Name < a[j].Name.Name }

// Moves all top-level functions to the end, sorted in alphabetical order.
// The "source file" is given as a string (rather than e.g. a filename).
func SortFunctions(src string) (string, error) {
	f, err := decorator.Parse(src)
	if err != nil {
		return "", err
	}

	funcs := make([]dst.FuncDecl, 0)
	rest := make([]dst.Decl, 0)

	declLen := len(f.Decls)
	for i := 0; i < declLen; i++ {
		it := f.Decls[i]

		if decl, ok := it.(*dst.FuncDecl); ok {
			funcs = append(funcs, *decl)
		} else {
			rest = append(rest, it)
		}
	}

	sort.Sort(ByName(funcs))
	f.Decls = rest
	for i := 0; i < len(funcs); i++ {
		it := funcs[i]
		f.Decls = append(f.Decls, dst.Decl(&it))
	}

	buf := new(bytes.Buffer)
	err = decorator.Fprint(buf, f)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func main() {
	f, err := decorator.Parse(src)
	if err != nil {
		log.Fatal(err)
	}

	// Print AST
	err = dst.Fprint(os.Stdout, f, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Convert AST back to source
	fmt.Printf("\n===ORIGINAL SOURCE===\n")
	err = decorator.Print(f)
	if err != nil {
		log.Fatal(err)
	}

	// Print sorted source
	fmt.Printf("\n===SORTED SOURCE===\n")
	sortedStr, err := SortFunctions(src)
	if err != nil {
		log.Fatal(err)
	}
	print(sortedStr)
}
