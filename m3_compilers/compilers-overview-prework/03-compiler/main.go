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
	if x == 0 {
		return y
	} else {
		return x
	}
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
	result, err := runVM(bytecode, 1, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
