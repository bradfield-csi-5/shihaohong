package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	argsLen := len(os.Args)

	if argsLen > 2 {
		fmt.Println("Usage: glox <inputFile>")
	} else if argsLen == 2 {
		l := &Lox{}
		err := l.runFile(os.Args[1])
		if err != nil {
			panic(err)
		}
	} else {
		l := &Lox{}
		err := l.runPrompt()
		if err != nil {
			panic(err)
		}
	}
}

type Lox struct {
	Scanner *Scanner
	Parser  *Parser
}

func (l *Lox) runFile(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	l.run(string(file))
	return nil
}

func (l *Lox) runPrompt() error {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		l.run(line)
		l.Scanner.hadError = false
	}
}

func (l *Lox) run(src string) {
	l.Scanner = NewScanner(src)
	tokens := l.Scanner.ScanTokens()
	for _, t := range tokens {
		fmt.Printf("token: %s\n", t.String())
	}

	l.Parser = NewParser(tokens)
	expr := l.Parser.parse()

	p := AstPrinter{}
	fmt.Println(p.print(expr))
}
