package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func generateBytecode(src string) ([]byte, error) {
	node, err := parse(src, "f")
	if err != nil {
		return nil, err
	}
	asm, err := compile(node)
	if err != nil {
		return nil, err
	}
	bytecode, err := assemble(asm)
	if err != nil {
		return nil, err
	}
	return bytecode, nil
}

func runVM(bytecode []byte, x, y byte) (byte, error) {
	// Set up the memory according to the expected layout
	memory := make([]byte, 256)
	copy(memory[instructionStart:], bytecode)
	memory[parameterStart] = x
	memory[parameterStart+1] = y

	// Actually run the VM
	err := execute(memory)
	if err != nil {
		return 0, err
	}

	// Return value is placed in memory location 0
	return memory[0], nil
}

const src string = `package f

func f(x, y byte) byte {
	var a byte
	var b byte
	var c byte
	a = 0
	b = 1
	for x > 0 {
		c = a + b
		a = b
		b = c
		x = x - 1
	}
	return b
}`

func main() {
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)

	bytecode, err := generateBytecode(src)
	if err != nil {
		log.Fatal(err)
	}
	result, err := runVM(bytecode, 9, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
