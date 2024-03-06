package parser

import (
	"testing"

	"github.com/abiriadev/clea/lexer"
	"github.com/stretchr/testify/assert"
)

func TestCleaLexer(t *testing.T) {
	parser := NewCleaParser("(+ 1 2)")

	assert := assert.New(t)

	assert.Equal([]Tree{TreeFromSlice([]Tree{
		TreeFromToken(lexer.Token{
			Type:  lexer.IdentType,
			Value: lexer.Ident("+"),
		}),
		TreeFromToken(lexer.Token{
			Type:  lexer.NumberType,
			Value: lexer.Number(1),
		}),
		TreeFromToken(lexer.Token{
			Type:  lexer.NumberType,
			Value: lexer.Number(2),
		}),
	})}, parser.Parse())
}
