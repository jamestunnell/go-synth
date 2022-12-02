package main

import (
	"fmt"
	"os"

	"github.com/jamestunnell/go-synth/slang/repl"
)

func main() {
	fmt.Println("go-synth wiring REPL")
	repl.Start(os.Stdin, os.Stdout)
}
