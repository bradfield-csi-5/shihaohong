package main

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strconv"
)

// Given an expression containing only int types, evaluate
// the expression and return the result.
//
// root of ParenExpr is always BinaryExpr
func Evaluate(expr ast.Expr) (int, error) {
	switch x := expr.(type) {
	case *ast.BasicLit:
		return handleBasicLit(x)
	case *ast.BinaryExpr:
		return handleBinaryExpr(x)
	case *ast.ParenExpr:
		return Evaluate(x.X)
	default:
		return 0, errors.New("undefined ast expression")
	}
}

func handleBinaryExpr(be *ast.BinaryExpr) (int, error) {
	xval, err := Evaluate(be.X)
	if err != nil {
		return 0, err
	}
	yval, err := Evaluate(be.Y)
	if err != nil {
		return 0, err
	}

	switch be.Op.String() {
	case "+":
		return xval + yval, nil
	case "-":
		return xval - yval, nil
	case "*":
		return xval * yval, nil
	case "/":
		return xval / yval, nil
	default:
		return 0, errors.New("undefined operation")
	}
}

func handleBasicLit(bl *ast.BasicLit) (int, error) {
	val, err := strconv.Atoi(bl.Value)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func main() {
	expr, err := parser.ParseExpr("((3 - 4)) * 2")
	if err != nil {
		log.Fatal(err)
	}
	fset := token.NewFileSet()
	err = ast.Print(fset, expr)
	if err != nil {
		log.Fatal(err)
	}
}
