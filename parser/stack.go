package parser

import (
	"github.com/abiriadev/clea/lexer"
)

type TokenStack struct {
	stack [][]Tree
}

func NewTokenStack() TokenStack {
	ts := TokenStack{
		stack: make([][]Tree, 0),
	}

	ts.Push()

	return ts
}

func (s *TokenStack) Push() {
	s.stack = append(s.stack, make([]Tree, 0))
}

func (s *TokenStack) Insert(token lexer.Token) {
	l := len(s.stack)
	s.stack[l-1] = append(s.stack[l-1], TreeFromToken(token))
}

func (s *TokenStack) Shift() {
	last := s.Pop()
	l := len(s.stack)

	s.stack[l-1] = append(s.stack[l-1], TreeFromSlice(last))
}

func (s *TokenStack) Pop() []Tree {
	l := len(s.stack)
	last := s.stack[l-1]
	s.stack = s.stack[:l-1]

	return last
}
