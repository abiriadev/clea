package parser

import (
	"github.com/abiriadev/clea/lexer"
	"github.com/samber/mo"
)

type Tree mo.Either[lexer.Token, []Tree]
