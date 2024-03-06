package lexer

import (
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
		scanner.SkipComments
	return CleaLexer{
		s,
	}
}

func (lexer *CleaLexer) Scan() TokenType {
	tt := lexer.scanner.Scan()

	return TokenTypeFromRune(tt)
}
