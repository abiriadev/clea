package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleaLexer(t *testing.T) {
	lexer := NewCleaLexer("(+ 1 2)")

	assert := assert.New(t)

	assert.Equal(LType, lexer.Scan().Type)
	assert.Equal(IdentType, lexer.Scan().Type)
	assert.Equal(NumberType, lexer.Scan().Type)
	assert.Equal(NumberType, lexer.Scan().Type)
	assert.Equal(RType, lexer.Scan().Type)
	assert.Equal(EofType, lexer.Scan().Type)
}

func TestShouldLexParenthesis(t *testing.T) {
	lexer := NewCleaLexer("(()())")

	assert := assert.New(t)

	assert.Equal(LType, lexer.Scan().Type)
	assert.Equal(LType, lexer.Scan().Type)
	assert.Equal(RType, lexer.Scan().Type)
	assert.Equal(LType, lexer.Scan().Type)
	assert.Equal(RType, lexer.Scan().Type)
	assert.Equal(RType, lexer.Scan().Type)
	assert.Equal(EofType, lexer.Scan().Type)
}

func TestShouldLexIdent(t *testing.T) {
	lexer := NewCleaLexer("abc + -")

	assert := assert.New(t)

	assert.Equal(IdentType, lexer.Scan().Type)
	assert.Equal(IdentType, lexer.Scan().Type)
	assert.Equal(IdentType, lexer.Scan().Type)
	assert.Equal(EofType, lexer.Scan().Type)
}

func TestShouldLexNumber(t *testing.T) {
	lexer := NewCleaLexer("123 4.56")

	assert := assert.New(t)

	assert.Equal(NumberType, lexer.Scan().Type)
	assert.Equal(NumberType, lexer.Scan().Type)
	assert.Equal(EofType, lexer.Scan().Type)
}

func TestShouldLexString(t *testing.T) {
	lexer := NewCleaLexer(`"hello, World!"`)

	assert := assert.New(t)

	assert.Equal(StringType, lexer.Scan().Type)
	assert.Equal(EofType, lexer.Scan().Type)
}

func TestShouldSkipComment(t *testing.T) {
	lexer := NewCleaLexer("// comment")

	assert := assert.New(t)

	assert.Equal(EofType, lexer.Scan().Type)
}
