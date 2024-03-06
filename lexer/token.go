package lexer

import (
	"text/scanner"
)

type TokenType int

const (
	EOF TokenType = iota
	L
	R
	Ident
	Number
	String

	Error
)

var tokenTypeMap = map[TokenType]string{
	EOF:    "EOF",
	L:      "(",
	R:      ")",
	Ident:  "Ident",
	Number: "Number",
	String: "String",
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
		return EOF
	case '(':
		return L
	case ')':
		return R
	case scanner.Ident:
	case '+':
	case '-':
		return Ident
	case scanner.Int:
	case scanner.Float:
		return Number
	case scanner.String:
		return String
	}

	return Error
}
