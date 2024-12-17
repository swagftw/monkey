package lexer

import "monkey/lexer/token"

type Lexer interface {
	NextToken() *token.Token
}

type lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) Lexer {
	l := &lexer{input: input}
	l.readChar()
	return l
}

func (l *lexer) NextToken() *token.Token {
	tok := new(token.Token)

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.TokenType = token.EOF
	default:
		return newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()

	return tok
}

// readChar reads the char on readPosition
// if readPosition is out of bounds then we assign ch to 0 byte
// else we assign ch to character at readPosition
// then we change position to readPosition and increment readPosition by 1 ie next char
func (l *lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) *token.Token {
	return &token.Token{
		TokenType: tokenType,
		Literal:   string(ch),
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
