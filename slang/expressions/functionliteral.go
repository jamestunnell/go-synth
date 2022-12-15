package expressions

import (
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

func (f *FunctionLiteral) Type() slang.ExprType {
	return slang.ExprFUNCTIONLITERAL
}

func (f *FunctionLiteral) Equal(other slang.Expression) bool {
	f2, ok := other.(*FunctionLiteral)
	if !ok {
		return false
	}

	return f.Body.Equal(f2.Body) &&
		slices.EqualFunc(f.Params, f2.Params, indentifiersEqual)

}

func indentifiersEqual(a, b *Identifier) bool {
	return a.Equal(b)
}
