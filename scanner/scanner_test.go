package scanner

import (
	"fmt"
	"testing"

	"github.com/jmptc/lox/token"
)

func TestAtTheEnd(t *testing.T) {
	scanner := NewScanner("Hello")

	if scanner.atTheEnd() {
		t.Fatalf("Expected: %t Got: %t. Current: %d Source Len: %d",
			false, scanner.atTheEnd(), scanner.current, len(scanner.source))
	}

	scanner.current = len(scanner.source)
	if !scanner.atTheEnd() {
		t.Fatalf("Expected: %t Got: %t. Current: %d Source Len: %d",
			true, scanner.atTheEnd(), scanner.current, len(scanner.source))
	}
}

func TestScanSingleCharacterTokens(t *testing.T) {
	source := "(){},.-+;/*"

	scanner := NewScanner(source)
	tokens := scanner.ScanTokens()

	expectedTokens := []token.Token{
		{token.LEFT_PAREN, "(", "("},
		{token.RIGHT_PAREN, ")", ")"},
		{token.LEFT_BRACE, "{", "{"},
		{token.RIGHT_BRACE, "}", "}"},
		{token.COMMA, ",", ","},
		{token.DOT, ".", "."},
		{token.MINUS, "-", "-"},
		{token.PLUS, "+", "+"},
		{token.SEMICOLON, ";", ";"},
		{token.SLASH, "/", "/"},
		{token.STAR, "*", "*"},
		{token.EOF, "", ""},
	}

	for i := 0; i < len(expectedTokens); i++ {
		if expectedTokens[i].Type != tokens[i].Type {
			t.Fatalf("%d Expected TokenType: %s Got: %s", i, expectedTokens[i].Type, tokens[i].Type)
		}
	}
}

func TestScanTwoCharacterTokens(t *testing.T) {
	source := "= == ! != < <= > >="

	scanner := NewScanner(source)
	tokens := scanner.ScanTokens()

	expectedTokens := []token.Token{
        {token.EQUAL, "=", "="},
        {token.EQUAL_EQUAL, "==", "=="},
        {token.BANG, "!", "!"},
        {token.BANG_EQUAL, "!=", "!="},
        {token.LESS, "<", "<"},
        {token.LESS_EQUAL, "<=", "<="},
        {token.GREATER, ">", ">"},
        {token.GREATER_EQUAL, ">=", ">="},
        {token.EOF, "", ""},
	}

    // fmt.Println(tokens)
	for i := 0; i < len(expectedTokens); i++ {
		if expectedTokens[i].Type != tokens[i].Type {
			t.Fatalf("%d Expected TokenType: %s Got: %s", i, expectedTokens[i].Type, tokens[i].Type)
		}
	}
}
