package lexer

import (
	"Monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	inputs := []string{
		`=+(){},;`,
		`let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
`,
		`!-/*67`,
		`3 < x > 9`,
		`if (5 < 10) {
  return true;
} else {
  return false;
}`,
	}

	results := [][]struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{
			{token.ASSIGN, "="},
			{token.PLUS, "+"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},
			{token.COMMA, ","},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		},
		{
			{token.LET, "let"},
			{token.IDENT, "five"},
			{token.ASSIGN, "="},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "ten"},
			{token.ASSIGN, "="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "add"},
			{token.ASSIGN, "="},
			{token.FUNCTION, "fn"},
			{token.LPAREN, "("},
			{token.IDENT, "x"},
			{token.COMMA, ","},
			{token.IDENT, "y"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.IDENT, "x"},
			{token.PLUS, "+"},
			{token.IDENT, "y"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "result"},
			{token.ASSIGN, "="},
			{token.IDENT, "add"},
			{token.LPAREN, "("},
			{token.IDENT, "five"},
			{token.COMMA, ","},
			{token.IDENT, "ten"},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		},
		{
			{token.BANG, "!"},
			{token.MINUS, "-"},
			{token.SLASH, "/"},
			{token.ASTERISK, "*"},
			{token.INT, "67"},
		},
		{
			{token.INT, "3"},
			{token.LT, "<"},
			{token.IDENT, "x"},
			{token.GT, ">"},
			{token.INT, "9"},
		},
		{
			{token.IF, "if"},
			{token.LPAREN, "("},
			{token.INT, "5"},
			{token.LT, "<"},
			{token.INT, "10"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RETURN, "return"},
			{token.TRUE, "true"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.ELSE, "else"},
			{token.LBRACE, "{"},
			{token.RETURN, "return"},
			{token.FALSE, "false"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
		},
	}

	for i, input := range inputs {
		t.Logf("When given the %d-th input", i)
		l := New(input)

		for j, tt := range results[i] {
			tok := l.NextToken()

			if tok.Type != tt.expectedType {
				t.Fatalf("%d-th token - tokentype wrong. expected=%q, got=%q",
					j, tt.expectedType, tok.Type)
			}

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("%d-th token - literal wrong. expected=%q, got=%q",
					j, tt.expectedType, tok.Literal)
			}
		}
	}
}
