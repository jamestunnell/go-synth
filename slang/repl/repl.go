package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jamestunnell/go-synth/slang/lexer"
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
		tokens := lexer.ScanString(line)

		for _, tok := range tokens {
			const strFmt = "{Type: %s, Value: %s, Loc: %s}\n"

			fmt.Fprintf(out, strFmt, tok.Info.Type(), tok.Info.Value(), tok.Location)
		}
	}
}
