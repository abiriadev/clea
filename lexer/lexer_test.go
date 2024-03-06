package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleaLexer(t *testing.T) {
	lexer := NewCleaLexer("(+ 1 2)")

	assert := assert.New(t)

	assert.Equal(lexer.Scan(), "(")
	assert.Equal(lexer.Scan(), "+")
	assert.Equal(lexer.Scan(), "1")
	assert.Equal(lexer.Scan(), "2")
	assert.Equal(lexer.Scan(), ")")
}
