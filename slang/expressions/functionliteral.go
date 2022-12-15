package expressions

import (
	"strings"

	"github.com/jamestunnell/go-synth/slang"
	"github.com/jamestunnell/go-synth/slang/statements"
	"golang.org/x/exp/slices"
)

type FunctionLiteral struct {
	Params []*Identifier
	Body   *statements.Block
}

func NewFunctionLiteral(
	params []*Identifier, body *statements.Block) *FunctionLiteral {
	return &FunctionLiteral{
		Params: params,
		Body:   body,
	}
}

func (af *FunctionLiteral) Type() slang.ExprType {
	return slang.ExprFUNCTIONLITERAL
}

func (af *FunctionLiteral) Equal(other slang.Expression) bool {
	af2, ok := other.(*FunctionLiteral)
	if !ok {
		return false
	}

	r := slices.CompareFunc(af.Params, af2.Params, cmpIdentifiersByName)
	if r != 0 {
		return false
	}

	return af.Body.Equal(af2.Body)
}

func cmpIdentifiersByName(a, b *Identifier) int {
	return strings.Compare(a.Name, b.Name)
}
