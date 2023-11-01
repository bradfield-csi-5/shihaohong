package main

type Expr interface {
	accept(Visitor) any
}

type BinaryExpr struct {
	left     Expr
	operator Token
	right    Expr
}

func (e *BinaryExpr) accept(v Visitor) any {
	return v.visitBinaryExpr(e)
}

type UnaryExpr struct {
	operator Token
	expr     Expr
}

func (e *UnaryExpr) accept(v Visitor) any {
	return v.visitUnaryExpr(e)
}

type LiteralExpr struct {
	value interface{}
}

func (e *LiteralExpr) accept(v Visitor) any {
	return v.visitLiteralExpr(e)
}

type GroupingExpr struct {
	expr Expr
}

func (e *GroupingExpr) accept(v Visitor) any {
	return v.visitGroupingExpr(e)
}

type Visitor interface {
	visitBinaryExpr(expr *BinaryExpr) any
	visitUnaryExpr(expr *UnaryExpr) any
	visitLiteralExpr(expr *LiteralExpr) any
	visitGroupingExpr(expr *GroupingExpr) any
}
