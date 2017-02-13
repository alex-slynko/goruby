package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/goruby/goruby/evaluator"
	"github.com/goruby/goruby/lexer"
	"github.com/goruby/goruby/parser"
)

const PROMPT = "girb:%03d> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	counter := 1
	for {
		fmt.Printf(PROMPT, counter)
		counter++
		scanned := scanner.Scan()
		if !scanned {
			fmt.Println()
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			fmt.Fprintf(out, "=> %s\n", evaluated.Inspect())
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	fmt.Println("Parser errors: ")
	for _, msg := range errors {
		fmt.Fprintf(out, "\t%s\n", msg)
	}
}
