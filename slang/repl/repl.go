package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/jamestunnell/go-synth/slang/interpreter"
	"github.com/jamestunnell/go-synth/slang/lexer"
	"github.com/jamestunnell/go-synth/slang/parser"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, prompt)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(strings.NewReader(line))
		p := parser.New(l)

		results := p.Run()
		if len(results.Errors) > 0 {
			for _, pErr := range results.Errors {
				l := pErr.Token.Location

				fmt.Fprintln(out, line)
				fmt.Fprintln(out, spaces(l.Column-1)+"^")

				fmt.Fprintf(out, "parser error at %s: %v\n", l, pErr.Error)
			}

			continue
		}

		for _, st := range results.Statements {
			obj := interpreter.EvalStatement(st)

			fmt.Fprintln(out, obj.Inspect())
		}
	}
}

func spaces(n int) string {
	return strings.Repeat(" ", n)
}
