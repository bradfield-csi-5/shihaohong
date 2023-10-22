package main

import (
	"errors"
	"fmt"
	"go/ast"
	"strconv"
)

// Given an expression containing only int types, evaluate
// the expression and return the result.
//
// root of ParenExpr is always BinaryExpr
func Evaluate(expr ast.Expr, data map[string]int) (string, error) {
	switch x := expr.(type) {
	case *ast.BasicLit:
		// fmt.Printf("basic lit\n")
		return handleBasicLit(x)
	case *ast.IndexExpr:
		// fmt.Printf("index expr\n")
		return "", nil
	case *ast.BinaryExpr:
		// fmt.Printf("binary expr\n")
		return handleBinaryExpr(x, data)
	case *ast.Ident:
		// fmt.Printf("ident\n")
		return handleIdentExpr(x, data)
	case *ast.ParenExpr:
		return Evaluate(x.X, data)
	default:
		return "", errors.New("undefined ast expression")
	}
}

func handleBinaryExpr(be *ast.BinaryExpr, data map[string]int) (string, error) {
	xval, err := Evaluate(be.X, data)
	if err != nil {
		return "", err
	}
	yval, err := Evaluate(be.Y, data)
	if err != nil {
		return "", err
	}

	switch be.Op.String() {
	case "+":
		return xval + yval + "add\n", nil
	case "-":
		return xval + yval + "sub\n", nil
	case "*":
		return xval + yval + "mul\n", nil
	case "/":
		return xval + yval + "div\n", nil
	default:
		return "", errors.New("undefined operation")
	}
}

func handleBasicLit(bl *ast.BasicLit) (string, error) {
	return "pushi " + bl.Value + "\n", nil
}

func handleIdentExpr(ie *ast.Ident, data map[string]int) (string, error) {
	val, ok := data[ie.Name]
	if !ok {
		return "", fmt.Errorf("undefined variable %s", ie.Name)
	}

	return "push " + strconv.Itoa(val) + "\n", nil
}

// Given an AST node corresponding to a function (guaranteed to be
// of the form `func f(x, y byte) byte`), compile it into assembly
// code.
//
// Recall from the README that the input parameters `x` and `y` should
// be read from memory addresses `1` and `2`, and the return value
// should be written to memory address `0`.
func compile(node *ast.FuncDecl) (string, error) {
	asm := ""
	// maps memory to var name
	data := make(map[string]int)
	data[node.Type.Params.List[0].Names[0].Name] = 1
	data[node.Type.Params.List[0].Names[1].Name] = 2

BodyLoop:
	for _, node := range node.Body.List {
		switch n := node.(type) {
		case *ast.ReturnStmt:
			// fmt.Printf("ASDF%+v\n", n.Results)

			val, err := Evaluate(n.Results[0], data)
			if err != nil {
				return "", err
			}
			asm += val
			asm += "pop 0\nhalt"
			break BodyLoop
		default:
			return "", errors.New("undefined ast node")
		}
	}

	fmt.Printf("end\n")
	fmt.Printf("asm: %v\n", asm)
	return asm, nil
}
