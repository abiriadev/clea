package lexer

type TokenType int

const (
	EOF TokenType = iota
	L
	R
	Ident
	Number
	String
)

var tokenTypeMap = map[TokenType]string{
	EOF:    "EOF",
	L:      "(",
	R:      ")",
	Ident:  "Ident",
	Number: "Number",
	String: "String",
}

func DebugTokenType(tokenType TokenType) string {
	if debug, found := tokenTypeMap[tokenType]; found {
		return debug
	}

	return ""
}
