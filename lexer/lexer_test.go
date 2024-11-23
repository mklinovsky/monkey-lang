package lexer

import (
	"testing"

	"github.com/mklinovsky/monkey-lang/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
	`

	expected := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LEFT_PARENTHESIS, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.LEFT_BRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RIGHT_BRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LEFT_PARENTHESIS, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RIGHT_PARENTHESIS, ")"},
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
