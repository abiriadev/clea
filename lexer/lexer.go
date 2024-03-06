package lexer

import (
	"strconv"
	"strings"
	"text/scanner"
)

type CleaLexer struct {
	scanner scanner.Scanner
}

func NewCleaLexer(source string) CleaLexer {
	var s scanner.Scanner
	s.Init(strings.NewReader(source))
	s.Mode = scanner.ScanIdents |
		scanner.ScanInts |
		scanner.ScanFloats |
		scanner.ScanStrings |
		scanner.ScanComments |
		scanner.SkipComments
	return CleaLexer{
		s,
	}
}

func (lexer *CleaLexer) Scan() Token {
	tt := lexer.scanner.Scan()

	Type := tokenTypeFromRune(tt)
	Value := tokenValueFromType(Type, lexer.scanner.String())

	return Token{
		Type,
		Value,
	}
}

func tokenTypeFromRune(rune rune) TokenType {
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

func tokenValueFromType(tokenType TokenType, text string) TokenValue {
	switch tokenType {
	case IdentType:
		return Ident(text)
	case NumberType:
		if f, err := strconv.ParseFloat(text, 64); err != nil {
			panic(err)
		} else {
			return Number(f)
		}
	case StringType:
		l := len(text)
		return String(text[1 : l-1])
	}

	// empty `TokenValue` for punctuation tokens
	return nil
}
