package main

import (
	"fmt"
	"strings"
)

type AstPrinter struct{}

// compile-time check to verify all methods are implemented
var _ Visitor = (*AstPrinter)(nil)

func (a *AstPrinter) print(expr Expr) string {
	str, ok := expr.accept(a).(string)
	if !ok {
		panic("expr did not evaluate to a string type")
	}
	return str
}

func (a *AstPrinter) visitBinaryExpr(expr *BinaryExpr) any {
	return a.parenthesize(expr.operator.lexeme, expr.left, expr.right)
}

func (a *AstPrinter) visitUnaryExpr(expr *UnaryExpr) any {
	return a.parenthesize(expr.operator.lexeme, expr.expr)

}

func (a *AstPrinter) visitLiteralExpr(expr *LiteralExpr) any {
	if expr.value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", expr.value)
}

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
