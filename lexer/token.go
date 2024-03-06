package lexer

import (
	"text/scanner"
)

type TokenType int

const (
	EofType TokenType = iota

	LType
	RType
	IdentType
	NumberType
	StringType

	ErrorType
)

var tokenTypeMap = map[TokenType]string{
	EofType:    "EOF",
	LType:      "(",
	RType:      ")",
	IdentType:  "Ident",
	NumberType: "Number",
	StringType: "String",
}

func DebugTokenType(tokenType TokenType) string {
	if debug, found := tokenTypeMap[tokenType]; found {
		return debug
	}

	return "Error"
}

func TokenTypeFromRune(rune rune) TokenType {
	switch rune {
	case scanner.EOF:
		return EofType
	case '(':
		return LType
	case ')':
		return RType
	case scanner.Ident, '+', '-':
		return IdentType
	case scanner.Int, scanner.Float:
		return NumberType
	case scanner.String:
		return StringType
	}

	return ErrorType
}
