package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleaLexer(t *testing.T) {
	lexer := NewCleaLexer("(+ 1 2)")

	assert := assert.New(t)

	assert.Equal(L, lexer.Scan())
	assert.Equal(Ident, lexer.Scan())
	assert.Equal(Number, lexer.Scan())
	assert.Equal(Number, lexer.Scan())
	assert.Equal(R, lexer.Scan())
	assert.Equal(EOF, lexer.Scan())
}

func TestShouldLexParenthesis(t *testing.T) {
	lexer := NewCleaLexer("(()())")

	assert := assert.New(t)

	assert.Equal(L, lexer.Scan())
	assert.Equal(L, lexer.Scan())
	assert.Equal(R, lexer.Scan())
	assert.Equal(L, lexer.Scan())
	assert.Equal(R, lexer.Scan())
	assert.Equal(R, lexer.Scan())
	assert.Equal(EOF, lexer.Scan())
}
