package lexer

import (
	"strings"
	"text/scanner"
)

type CleaLexer struct {
	scanner scanner.Scanner
}

func NewCleaLexer(source string) CleaLexer {
	var scanner scanner.Scanner
	scanner.Init(strings.NewReader(source))
	return CleaLexer{
		scanner,
	}
}

func (lexer *CleaLexer) Scan() string {
	lexer.scanner.Scan()
	return lexer.scanner.TokenText()
}