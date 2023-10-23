package main

import (
	"errors"
	"fmt"
	"go/ast"
	"strconv"
)

// TODO: pass around and initialize rather than leaving the counter as a global?
var labelCounter = 0
var labelTag = "label_"
var labelStartTag = "label_start_"
var labelEndTag = "label_end_"

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
	case "<":
		return fmt.Sprintf("%s%slt\n", xval, yval), nil
	case ">":
		return fmt.Sprintf("%s%sgt\n", xval, yval), nil
	case "==":
		return fmt.Sprintf("%s%seq\n", xval, yval), nil
	default:
		return "", fmt.Errorf("undefined operation %s", be.Op.String())
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
	data := make(map[string]int) // maps memory to var name
	labelCounter = 0             // reset counter

	data[node.Type.Params.List[0].Names[0].Name] = 1
	data[node.Type.Params.List[0].Names[1].Name] = 2

	res, err := processStmt(node.Body.List, data)
	if err != nil {
		return "", err
	}
	asm += res

	fmt.Printf("end\n")
	fmt.Printf("asm: %v\n", asm)
	return asm, nil
}

func processStmt(stmtList []ast.Stmt, data map[string]int) (string, error) {
	asm := ""

	for _, node := range stmtList {
		switch n := node.(type) {
		case *ast.AssignStmt:
			if lhs, ok := n.Lhs[0].(*ast.Ident); ok {
				lhsLoc, ok := data[lhs.Name]
				if !ok {
					dataLen := len(data)
					if len(data) >= 8 {
						return "", errors.New("not enough data space")
					}
					data[lhs.Name] = dataLen + 1
					lhsLoc = dataLen + 1
				}

				rhsRes, err := Evaluate(n.Rhs[0], data)
				if err != nil {
					return "", err
				}

				asm += rhsRes
				asm += "pop " + strconv.Itoa(lhsLoc) + "\n"
			} else {
				return "", errors.New("unexpected lhs")
			}
		case *ast.ReturnStmt:
			// fmt.Printf("ASDF%+v\n", n.Results)

			val, err := Evaluate(n.Results[0], data)
			if err != nil {
				return "", err
			}
			asm += val
			asm += "pop 0\nhalt\n"
			return asm, nil
		case *ast.IfStmt:
			res, err := Evaluate(n.Cond, data)
			if err != nil {
				return "", err
			}

			res += fmt.Sprintf("jeqz %s%d\n", labelTag, labelCounter)
			ifStmt, err := processStmt(n.Body.List, data)
			if err != nil {
				return "", err
			}
			res += ifStmt
			res += fmt.Sprintf("jump %s%d\n", labelEndTag, labelCounter)
			res += fmt.Sprintf("label %s%d\n", labelTag, labelCounter)

			elseStmtStruct := n.Else.(*ast.BlockStmt)
			elseStmt, err := processStmt(elseStmtStruct.List, data)
			if err != nil {
				return "", err
			}
			res += elseStmt
			res += fmt.Sprintf("label %s%d\n", labelEndTag, labelCounter)

			labelCounter++
			asm += res
		case *ast.DeclStmt:
			genDecl := n.Decl.(*ast.GenDecl)
			valueSpec := genDecl.Specs[0].(*ast.ValueSpec)
			declName := valueSpec.Names[0].Name
			_, ok := data[declName]
			if !ok {
				dataLen := len(data)
				if len(data) >= 8 {
					return "", errors.New("not enough data space")
				}
				data[declName] = dataLen + 1
			}
		case *ast.ForStmt:
			res := ""
			res += fmt.Sprintf("label %s%d\n", labelStartTag, labelCounter)
			condVal, err := Evaluate(n.Cond, data)
			if err != nil {
				return "", err
			}
			res += condVal

			res += fmt.Sprintf("jeqz %s%d\n", labelEndTag, labelCounter)
			stmt, err := processStmt(n.Body.List, data)
			if err != nil {
				return "", err
			}
			res += stmt
			res += fmt.Sprintf("jump %s%d\n", labelStartTag, labelCounter)
			res += fmt.Sprintf("label %s%d\n", labelEndTag, labelCounter)

			labelCounter++
			asm += res
		default:
			return "", errors.New("undefined ast node")
		}
	}
	return asm, nil
}
