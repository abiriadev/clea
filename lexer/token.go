package lexer

type TokenType int

const (
	EofType TokenType = iota

	LType
	RType
	IdentType
	NumberType
	StringType

	ErrorType
)

var tokenTypeMap = map[TokenType]string{
	EofType:    "EOF",
	LType:      "(",
	RType:      ")",
	IdentType:  "Ident",
	NumberType: "Number",
	StringType: "String",
}

func DebugTokenType(tokenType TokenType) string {
	if debug, found := tokenTypeMap[tokenType]; found {
		return debug
	}

	return "Error"
}

type Token struct {
	Type  TokenType
	Value TokenValue
}

type TokenValue interface {
	IsTokenValue()
}

type Ident string

func (_ Ident) IsTokenValue() {}

type Number float64

func (_ Number) IsTokenValue() {}

type String string

func (_ String) IsTokenValue() {}
