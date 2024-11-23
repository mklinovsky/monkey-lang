package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL           = "ILLEGAL"
	EOF               = "EOF"
	IDENTIFIER        = "IDENTIFIER"
	INT               = "INT"
	ASSIGN            = "="
	PLUS              = "+"
	COMMA             = ","
	SEMICOLON         = ";"
	LEFT_PARENTHESIS  = "("
	RIGHT_PARENTHESIS = ")"
	LEFT_BRACE        = "{"
	RIGHT_BRACE       = "}"
	FUNCTION          = "FUNCTION"
	LET               = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func GetKeywordOrIdentifier(word string) TokenType {
	if keyword, ok := keywords[word]; ok {
		return keyword
	}

	return IDENTIFIER
}
