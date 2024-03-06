package parser

import "github.com/abiriadev/clea/lexer"

type TokenStack struct {
	stack []lexer.Token
}

func NewTokenStack() TokenStack {
	return TokenStack{
		stack: make([]lexer.Token, 0),
	}
}
