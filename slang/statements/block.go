package statements

import "github.com/jamestunnell/go-synth/slang"

type Block struct {
	Statements []slang.Statement
}

func NewBlock(stmts ...slang.Statement) *Block {
	return &Block{Statements: stmts}
}

func (b *Block) Type() slang.StatementType {
	return slang.StatementEXPRESSION
}

func (b *Block) Equal(other slang.Statement) bool {
	b2, ok := other.(*Block)
	if !ok {
		return false
	}

	if len(b.Statements) != len(b2.Statements) {
		return false
	}

	for i, s := range b.Statements {
		if !b2.Statements[i].Equal(s) {
			return false
		}
	}

	return true
}
