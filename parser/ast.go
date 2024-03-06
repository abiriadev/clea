package parser

import (
	"github.com/abiriadev/clea/lexer"
	"github.com/samber/mo"
)

type Tree mo.Either[lexer.Token, []Tree]

func TreeFromToken(token lexer.Token) Tree {
	return Tree(mo.Left[lexer.Token, []Tree](token))
}

func TreeFromSlice(slice []Tree) Tree {
	return Tree(mo.Right[lexer.Token](slice))
}
