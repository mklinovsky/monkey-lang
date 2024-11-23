package lexer

import "github.com/mklinovsky/monkey-lang/token"

type Lexer struct {
	input        string
	currentChar  byte
	position     int
	readPosition int
}

func CreateLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()

	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.currentChar = 0
	} else {
		lexer.currentChar = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func createToken(tokenType token.TokenType, char byte) token.Token {
	if tokenType == token.EOF {
		return token.Token{Type: token.EOF, Literal: ""}
	}

	return token.Token{Type: tokenType, Literal: string(char)}
}

func (lexer *Lexer) NextToken() token.Token {
	var currentToken token.Token

	switch lexer.currentChar {
	case '=':
		currentToken = createToken(token.ASSIGN, lexer.currentChar)
	case ';':
		currentToken = createToken(token.SEMICOLON, lexer.currentChar)
	case '(':
		currentToken = createToken(token.LEFT_PARENTHESIS, lexer.currentChar)
	case ')':
		currentToken = createToken(token.RIGHT_PARENTHESIS, lexer.currentChar)
	case ',':
		currentToken = createToken(token.COMMA, lexer.currentChar)
	case '+':
		currentToken = createToken(token.PLUS, lexer.currentChar)
	case '{':
		currentToken = createToken(token.LEFT_BRACE, lexer.currentChar)
	case '}':
		currentToken = createToken(token.RIGHT_BRACE, lexer.currentChar)
	case 0:
		currentToken = createToken(token.EOF, lexer.currentChar)
	}

	lexer.readChar()
	return currentToken
}
