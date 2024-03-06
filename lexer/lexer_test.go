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

func TestShouldLexIdent(t *testing.T) {
	lexer := NewCleaLexer("abc + -")

	assert := assert.New(t)

	assert.Equal(Ident, lexer.Scan())
	assert.Equal(Ident, lexer.Scan())
	assert.Equal(Ident, lexer.Scan())
	assert.Equal(EOF, lexer.Scan())
}

func TestShouldLexNumber(t *testing.T) {
	lexer := NewCleaLexer("123 4.56")

	assert := assert.New(t)

	assert.Equal(Number, lexer.Scan())
	assert.Equal(Number, lexer.Scan())
	assert.Equal(EOF, lexer.Scan())
}

func TestShouldLexString(t *testing.T) {
	lexer := NewCleaLexer(`"hello, World!"`)

	assert := assert.New(t)

	assert.Equal(String, lexer.Scan())
	assert.Equal(EOF, lexer.Scan())
}
