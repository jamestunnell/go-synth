package parser

import (
	"fmt"
	"strconv"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/expressions"
)

func (p *Parser) parseInteger(str string) (slang.Expression, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse '%s' as int: %w", str, err)
	}

	return expressions.NewInteger(i), nil
}
