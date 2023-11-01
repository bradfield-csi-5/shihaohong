package main

import (
	"fmt"
	"strings"
)

type AstPrinter struct{}

func (a *AstPrinter) print(expr Expr) string {
	str, ok := expr.accept(a).(string)
	if !ok {
		panic("expr did not evaluate to a string type")
	}
	return str
}

//lint:ignore U1000 the visitor pattern is making this incorrectly report as unused
func (a *AstPrinter) visitBinaryExpr(expr *BinaryExpr) any {
	return a.parenthesize(expr.operator.lexeme, expr.left, expr.right)
}

//lint:ignore U1000 the visitor pattern is making this incorrectly report as unused
func (a *AstPrinter) visitUnaryExpr(expr *UnaryExpr) any {
	return a.parenthesize(expr.operator.lexeme, expr.expr)

}

//lint:ignore U1000 the visitor pattern is making this incorrectly report as unused
func (a *AstPrinter) visitLiteralExpr(expr *LiteralExpr) any {
	if expr.value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", expr.value)
}

//lint:ignore U1000 the visitor pattern is making this incorrectly report as unused
func (a *AstPrinter) visitGroupingExpr(expr *GroupingExpr) any {
	return a.parenthesize("group", expr.expr)
}

func (a *AstPrinter) parenthesize(name string, exprs ...Expr) any {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(name)

	for _, expr := range exprs {
		sb.WriteString(" ")
		str, ok := expr.accept(a).(string)
		if !ok {
			panic("expr did not evaluate to a string type")
		}
		sb.WriteString(str)
	}

	sb.WriteString(")")
	return sb.String()
}
