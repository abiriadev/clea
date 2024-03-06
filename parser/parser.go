package parser

import "github.com/abiriadev/clea/lexer"

type CleaParser struct {
	lexer lexer.CleaLexer
	stack TokenStack
}

func NewCleaParser(source string) CleaParser {
	return CleaParser{
		lexer: lexer.NewCleaLexer(source),
		stack: NewTokenStack(),
	}
}

func (parser *CleaParser) Parse() []Tree {
	for tok := parser.lexer.Scan(); tok.Type != lexer.EofType; tok = parser.lexer.Scan() {
		switch tok.Type {
		case lexer.LType:
			parser.stack.Push()
		case lexer.RType:
			parser.stack.Shift()
		default:
			parser.stack.Insert(tok)
		}
	}

	return parser.stack.Pop()
}
