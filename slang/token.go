package slang

import "fmt"

type TokenType int

type TokenInfo interface {
	Type() TokenType
	Value() string
}

type SourceLocation struct {
	Line, Column int
}

type Token struct {
	Info     TokenInfo
	Location SourceLocation
}

func NewToken(info TokenInfo, loc SourceLocation) *Token {
	return &Token{
		Info:     info,
		Location: loc,
	}
}

func NewLoc(line, col int) SourceLocation {
	return SourceLocation{Line: line, Column: col}
}

func (loc SourceLocation) String() string {
	return fmt.Sprintf("(line: %d, col: %d)", loc.Line, loc.Column)
}

const (
	TokenASSIGN TokenType = iota
	TokenBANG
	TokenCOMMA
	TokenDOT
	TokenIDENT
	TokenIF
	TokenELSE
	TokenEOF
	TokenEQUAL
	TokenFALSE
	TokenFLOAT
	TokenFUNC
	TokenGREATER
	TokenGREATEREQUAL
	TokenILLEGAL
	TokenINT
	TokenLBRACE
	TokenLESS
	TokenLESSEQUAL
	TokenLINE
	TokenLPAREN
	TokenMINUS
	TokenMINUSEQUAL
	TokenMINUSMINUS
	TokenNOTEQUAL
	TokenPLUS
	TokenPLUSEQUAL
	TokenPLUSPLUS
	TokenRBRACE
	TokenRETURN
	TokenRPAREN
	TokenSEMICOLON
	TokenSLASH
	TokenSLASHEQUAL
	TokenSTAR
	TokenSTAREQUAL
	TokenTRUE
)

func (tt TokenType) String() string {
	var str string

	switch tt {
	case TokenASSIGN:
		str = "ASSIGN"
	case TokenBANG:
		str = "BANG"
	case TokenCOMMA:
		str = "COMMA"
	case TokenDOT:
		str = "DOT"
	case TokenIDENT:
		str = "IDENT"
	case TokenIF:
		str = "IF"
	case TokenELSE:
		str = "ELSE"
	case TokenEOF:
		str = "EOF"
	case TokenEQUAL:
		str = "EQUAL"
	case TokenFALSE:
		str = "FALSE"
	case TokenFLOAT:
		str = "FLOAT"
	case TokenFUNC:
		str = "FUNC"
	case TokenGREATER:
		str = "GREATER"
	case TokenGREATEREQUAL:
		str = "GREATEREQUAL"
	case TokenILLEGAL:
		str = "ILLEGAL"
	case TokenINT:
		str = "INT"
	case TokenLBRACE:
		str = "LBRACE"
	case TokenLESS:
		str = "LESS"
	case TokenLESSEQUAL:
		str = "LESSEQUAL"
	case TokenLINE:
		str = "LINE"
	case TokenLPAREN:
		str = "LPAREN"
	case TokenMINUS:
		str = "MINUS"
	case TokenMINUSEQUAL:
		str = "MINUSEQUAL"
	case TokenMINUSMINUS:
		str = "MINUSMINUS"
	case TokenNOTEQUAL:
		str = "NOTEQUAL"
	case TokenPLUS:
		str = "PLUS"
	case TokenPLUSEQUAL:
		str = "PLUSEQUAL"
	case TokenPLUSPLUS:
		str = "PLUSPLUS"
	case TokenRBRACE:
		str = "RBRACE"
	case TokenRETURN:
		str = "RETURN"
	case TokenRPAREN:
		str = "RPAREN"
	case TokenSEMICOLON:
		str = "SEMICOLON"
	case TokenSLASH:
		str = "SLASH"
	case TokenSLASHEQUAL:
		str = "SLASHEQUAL"
	case TokenSTAR:
		str = "STAR"
	case TokenSTAREQUAL:
		str = "STAREQUAL"
	case TokenTRUE:
		str = "TRUE"
	}

	return str
}
