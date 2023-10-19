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
func Evaluate(expr ast.Expr) (int, error) {
	switch x := expr.(type) {
	case *ast.BasicLit:
		return handleBasicLit(x)
	case *ast.BinaryExpr:
		// root of ParenExpr is always BinaryExpr
		return handleBinaryExpr(x)
	default:
		return 0, errors.New("undefined ast expression")
	}
}

func handleBinaryExpr(be *ast.BinaryExpr) (int, error) {
	xval, yval := 0, 0
	switch xExpr := be.X.(type) {
	case *ast.BasicLit:
		var err error
		xval, err = handleBasicLit(xExpr)
		if err != nil {
			return 0, err
		}
	case *ast.BinaryExpr:
		var err error
		xval, err = handleBinaryExpr(xExpr)
		if err != nil {
			return 0, err
		}
	case *ast.ParenExpr:
		var err error
		xval, err = handleParen(xExpr.X)
		print(xval)
		if err != nil {
			return 0, err
		}
	}

	switch yExpr := be.Y.(type) {
	case *ast.BasicLit:
		var err error
		yval, err = handleBasicLit(yExpr)
		if err != nil {
			return 0, err
		}
	case *ast.BinaryExpr:
		var err error
		yval, err = handleBinaryExpr(yExpr)
		if err != nil {
			return 0, err
		}
	case *ast.ParenExpr:
		var err error
		yval, err = handleParen(yExpr.X)
		if err != nil {
			return 0, err
		}
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

// X can be a binary expr, basic lit, or another paren
func handleParen(expr ast.Expr) (int, error) {
	switch x := expr.(type) {
	case *ast.BasicLit:
		return handleBasicLit(x)
	case *ast.BinaryExpr:
		return handleBinaryExpr(x)
	case *ast.ParenExpr:
		return handleParen(x.X)
	default:
		return 0, errors.New("undefined ast expression")
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
