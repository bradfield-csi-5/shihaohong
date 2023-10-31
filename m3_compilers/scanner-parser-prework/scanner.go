package main

import (
	"fmt"
	"strconv"
)

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
}

func NewScanner(input string) *Scanner {
	return &Scanner{
		source:   input,
		tokens:   make([]Token, 0),
		start:    0,
		current:  0,
		line:     1,
		hadError: false,
	}
}

type Scanner struct {
	source   string
	tokens   []Token
	start    int
	current  int
	line     int
	hadError bool
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, Token{EOF, "", nil, s.line})
	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
	case '/':
		if s.match('/') {
			// A comment goes until the end of the line.
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
	case ' ':
	case '\r':
	case '\t':
		// Ignore whitespace.
	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		if s.isDigit(c) {
			s.number()
		} else if s.isAlpha(c) {
			s.identifier()
		} else {
			s.error(s.line, "Unterminated string.")
		}
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

// assume ASCII bytes in utf-8 for now, here and elsewhere
func (s *Scanner) advance() uint8 {
	c := s.source[s.current]
	s.current++
	return c
}

func (s *Scanner) peek() uint8 {
	if s.isAtEnd() {
		// src: https://go.dev/ref/spec#String_literals
		// The three-digit octal (\nnn) and two-digit hexadecimal (\xnn)
		// escapes represent individual bytes of the resulting string
		return "\x00"[0]
	}
	return s.source[s.current]
}

func (s *Scanner) peekNext() uint8 {
	if s.current+1 >= len(s.source) {
		// src: https://go.dev/ref/spec#String_literals
		// The three-digit octal (\nnn) and two-digit hexadecimal (\xnn)
		// escapes represent individual bytes of the resulting string
		return "\x00"[0]
	}
	return s.source[s.current+1]
}

func (s *Scanner) addToken(tokenType TokenType) {
	s.addTokenWithLiteral(tokenType, nil)
}

func (s *Scanner) addTokenWithLiteral(tokenType TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	var newToken = Token{
		tokenType: tokenType,
		lexeme:    text,
		literal:   literal,
		line:      s.line,
	}
	s.tokens = append(s.tokens, newToken)
}

func (s *Scanner) match(expected uint8) bool {
	if s.isAtEnd() {
		return false
	} else if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.error(s.line, "Unterminated string.")
		return
	}

	// closing '"'
	s.advance()

	// Trim out the quotes
	value := s.source[s.start+1 : s.current-1]
	s.addTokenWithLiteral(STRING, value)
}

func (s *Scanner) isDigit(c uint8) bool {
	return '0' <= c && c <= '9'
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		// Consume '.'
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}
	}

	// TODO: handle float err?
	float, _ := strconv.ParseFloat(s.source[s.start:s.current], 64)
	s.addTokenWithLiteral(NUMBER, float)
}

func (s *Scanner) isAlpha(c uint8) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func (s *Scanner) isAlphaNumeric(c uint8) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

func (s *Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}

	value := s.source[s.start:s.current]

	tokenType, found := keywords[value]
	if !found {
		tokenType = IDENTIFIER
	}

	s.addToken(tokenType)
}

func (s *Scanner) error(line int, message string) {
	s.report(line, "", message)
}

func (s *Scanner) report(line int, where string, message string) {
	fmt.Printf("<line %d> Error%s: %s\n", line, where, message)
	s.hadError = false
}
