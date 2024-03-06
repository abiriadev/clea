package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleaLexer(t *testing.T) {
	lexer := NewCleaLexer("(+ 1 2)")

	assert := assert.New(t)

	assert.Equal(LType, lexer.Scan())
	assert.Equal(IdentType, lexer.Scan())
	assert.Equal(NumberType, lexer.Scan())
	assert.Equal(NumberType, lexer.Scan())
	assert.Equal(RType, lexer.Scan())
	assert.Equal(EofType, lexer.Scan())
}

func TestShouldLexParenthesis(t *testing.T) {
	lexer := NewCleaLexer("(()())")

	assert := assert.New(t)

	assert.Equal(LType, lexer.Scan())
	assert.Equal(LType, lexer.Scan())
	assert.Equal(RType, lexer.Scan())
	assert.Equal(LType, lexer.Scan())
	assert.Equal(RType, lexer.Scan())
	assert.Equal(RType, lexer.Scan())
	assert.Equal(EofType, lexer.Scan())
}

func TestShouldLexIdent(t *testing.T) {
	lexer := NewCleaLexer("abc + -")

	assert := assert.New(t)

	assert.Equal(IdentType, lexer.Scan())
	assert.Equal(IdentType, lexer.Scan())
	assert.Equal(IdentType, lexer.Scan())
	assert.Equal(EofType, lexer.Scan())
}

func TestShouldLexNumber(t *testing.T) {
	lexer := NewCleaLexer("123 4.56")

	assert := assert.New(t)

	assert.Equal(NumberType, lexer.Scan())
	assert.Equal(NumberType, lexer.Scan())
	assert.Equal(EofType, lexer.Scan())
}

func TestShouldLexString(t *testing.T) {
	lexer := NewCleaLexer(`"hello, World!"`)

	assert := assert.New(t)

	assert.Equal(StringType, lexer.Scan())
	assert.Equal(EofType, lexer.Scan())
}

func TestShouldSkipComment(t *testing.T) {
	lexer := NewCleaLexer("// comment")

	assert := assert.New(t)

	assert.Equal(EofType, lexer.Scan())
}
