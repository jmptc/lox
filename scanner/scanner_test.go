package scanner

import (
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

func TestScanTokens(t *testing.T) {
    source := "(){},."

    scanner := NewScanner(source)
    tokens := scanner.ScanTokens()

    expectedTokens := []token.Token{
        token.Token{token.LEFT_PAREN, "(", "("},
        token.Token{token.RIGHT_PAREN, ")", ")"},
        token.Token{token.LEFT_BRACE, "{", "{"},
        token.Token{token.RIGHT_BRACE, "}", "}"},
        token.Token{token.COMMA, ",", ","},
        token.Token{token.DOT, ".", "."},
        token.Token{token.EOF, "", ""},
    }

    for i := 0; i < len(expectedTokens); i++ {
        if expectedTokens[i].Type != tokens[i].Type {
            t.Fatalf("%d Expected TokenType: %s Got: %s", i, expectedTokens[i].Type, tokens[i].Type)
        }
    }
}
