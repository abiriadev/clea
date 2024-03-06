package parser

import (
	"github.com/abiriadev/clea/lexer"
	"github.com/samber/mo"
)

type TokenStack struct {
	stack []Tree
}

func NewTokenStack() TokenStack {
	return TokenStack{
		stack: make([]Tree, 0),
	}
}

func (s *TokenStack) Push() {
	s.stack = append(s.stack, Tree(mo.Right[lexer.Token](make([]Tree, 0))))
}

func (s *TokenStack) Insert(token lexer.Token) {
	s.stack = append(s.stack, Tree(mo.Left[lexer.Token, []Tree](token)))
}
