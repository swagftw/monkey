package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// identifier + literals
	IDENTIFIER TokenType = "IDENT"
	INT        TokenType = "INT"

	ASSIGN TokenType = "="
	PLUS   TokenType = "+"

	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	// keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
)

type Token struct {
	TokenType TokenType
	Literal   string
}
