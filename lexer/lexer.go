package lexer

import (
	"K/token"
)

type Lexer struct {
	input        string
	position     int  //current position in input (points to current char)
	readPosition int  //current reading position in input (after current char)
	ch           byte //current char under examination
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]

	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = NewToken(token.ASSIGN, l.ch)
	case ';':
		tok = NewToken(token.SEMICOLON, l.ch)
	case '(':
		tok = NewToken(token.LEFT_BRACKET, l.ch)
	case ')':
		tok = NewToken(token.RIGHT_BRACKET, l.ch)
	case ',':
		tok = NewToken(token.COMMA, l.ch)
	case '+':
		tok = NewToken(token.PLUS, l.ch)
	case '-':
		tok = NewToken(token.MINUS, l.ch)
	case '{':
		tok = NewToken(token.LEFT_CURL_BRACKET, l.ch)
	case '}':
		tok = NewToken(token.RIGHT_CURL_BRACKET, l.ch)
	case '!':
		tok = NewToken(token.EXCLAMATION, l.ch)
	case '*':
		tok = NewToken(token.ASTERISK, l.ch)
	case '>':
		tok = NewToken(token.GT, l.ch)
	case '<':
		tok = NewToken(token.LT, l.ch)
	case '/':
		tok = NewToken(token.SLASH, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok

		} else {
			tok = NewToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func NewToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}
