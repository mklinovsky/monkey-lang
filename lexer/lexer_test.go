package lexer

import (
	"testing"

	"github.com/mklinovsky/monkey-lang/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	expected := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LEFT_PARENTHESIS, "("},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.LEFT_BRACE, "{"},
		{token.RIGHT_BRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := CreateLexer(input)

	for i, expectedToken := range expected {
		token := lexer.NextToken()

		if token.Type != expectedToken.expectedType {
			t.Fatalf("expected[%d] - tokentType wrong. expected=%q, got=%q",
				i, expectedToken.expectedType, token.Type)
		}

		if token.Literal != expectedToken.expectedLiteral {
			t.Fatalf("expected[%d] - literal wrong. expected=%q, got=%q",
				i, expectedToken.expectedLiteral, token.Literal)
		}
	}
}
