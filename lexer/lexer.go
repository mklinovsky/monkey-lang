package lexer

import (
	"github.com/mklinovsky/monkey-lang/token"
)

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

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && 'Z' <= char || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position

	for isLetter(lexer.currentChar) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.currentChar) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.currentChar == ' ' || lexer.currentChar == '\t' || lexer.currentChar == '\n' || lexer.currentChar == '\r' {
		lexer.readChar()
	}
}

func (lexer *Lexer) NextToken() token.Token {
	var currentToken token.Token

	lexer.skipWhitespace()

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
	default:
		if isLetter(lexer.currentChar) {
			currentToken.Literal = lexer.readIdentifier()
			currentToken.Type = token.GetKeywordOrIdentifier(currentToken.Literal)

			return currentToken
		} else if isDigit(lexer.currentChar) {
			currentToken.Type = token.INT
			currentToken.Literal = lexer.readNumber()

			return currentToken
		} else {
			currentToken = createToken(token.ILLEGAL, lexer.currentChar)
		}
	}

	lexer.readChar()
	return currentToken
}
