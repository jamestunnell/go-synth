package parser

import (
	"fmt"
	"strconv"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
)

func (p *Parser) parseFloat(str string) (slang.Expression, error) {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse '%s' as float: %w", str, err)
	}

	return expressions.NewFloat(f), nil
}
