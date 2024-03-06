package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleaLexer(t *testing.T) {
	lexer := NewCleaLexer("(+ 1 2)")

	assert := assert.New(t)

	assert.Equal(lexer.Scan(), L)
	assert.Equal(lexer.Scan(), Ident)
	assert.Equal(lexer.Scan(), Number)
	assert.Equal(lexer.Scan(), Number)
	assert.Equal(lexer.Scan(), R)
}
