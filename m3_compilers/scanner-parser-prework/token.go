package main

import "fmt"

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   interface{}
	line      int // [location]
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s %v", tokens[t.tokenType], t.lexeme, t.literal)
}
