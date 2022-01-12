package lexer

import (
	"testing"

	"kaizen_lang/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, " "},
	}

	l := New(input)

	for i, tt := range tests {
		token := l.NextToken()

		if token.Type != tt.expectedType {
			t.Fatalf(
				"Tests[%d] - Incorrect tokenType.\n\t\tExpected: %q\n\t\tGot: %q", i, tt.expectedType, token.Type,
			)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf(
				"Tests[%d] - Incorrect tokenLiteral.\n\t\tExpected: %q\n\t\tGot: %q", i, tt.expectedLiteral, token.Literal,
			)
		}
	}
}
