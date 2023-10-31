package main

type Expr interface {
	accept(Visitor[any]) any
}

// BinaryExpr implements rule: `binary -> expression operator expression;`
type BinaryExpr struct {
	left     Expr
	operator Token
	right    Expr
}

func (e *BinaryExpr) accept(v Visitor[any]) any {
	return v.visitBinaryExpr(e)
}

// UnaryExpr implements rule: `unary -> ("-" | "!") expression;`
type UnaryExpr struct {
	operator Token
	expr     Expr
}

func (e *UnaryExpr) accept(v Visitor[any]) any {
	return v.visitUnaryExpr(e)
}

// LiteralExpr implements rule: `literal -> NUMBER | STRING | "true" | "false" | "nil;`
type LiteralExpr struct {
	value interface{}
}

func (e *LiteralExpr) accept(v Visitor[any]) any {
	return v.visitLiteralExpr(e)
}

// GroupingExpr implements rule: `grouping -> "(" expression ")";`
type GroupingExpr struct {
	expr Expr
}

func (e *GroupingExpr) accept(v Visitor[any]) any {
	return v.visitGroupingExpr(e)
}

type Visitor[T any] interface {
	visitBinaryExpr(expr *BinaryExpr) T
	visitUnaryExpr(expr *UnaryExpr) T
	visitLiteralExpr(expr *LiteralExpr) T
	visitGroupingExpr(expr *GroupingExpr) T
}
