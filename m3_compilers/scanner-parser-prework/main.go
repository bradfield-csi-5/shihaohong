package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// TODO: implement runFile mode
	l := &Lox{}
	l.runPrompt()
}

type Lox struct {
	hadError bool
	Scanner  *Scanner
}

func (l *Lox) runPrompt() error {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		l.run(line)
		l.hadError = false
	}
}

func (l *Lox) run(src string) error {
	// TODO: implement scanner
	l.Scanner = NewScanner(src)
	tokens := l.Scanner.ScanTokens()
	for _, t := range tokens {
		fmt.Printf("token: %s\n", t.String())
	}

	return nil
	// TODO: implement parser
}
