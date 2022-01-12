package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers and literals
	IDENT = "IDENT"
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS = "+"

	//delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	KEYWORDS = "FUNCTION"
	LET = "LET"
)