package parser

import "github.com/jamestunnell/go-synth/slang"

type ParseErr struct {
	Error   error
	Token   *slang.Token
	Context *ParseContext
}

func NewParseError(err error, tok *slang.Token, ctx *ParseContext) *ParseErr {
	return &ParseErr{
		Error:   err,
		Token:   tok,
		Context: ctx,
	}
}
