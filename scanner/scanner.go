package scanner

import (
	"github.com/jmptc/lox/token"
)

type Scanner struct {
	source  string
	tokens  []token.Token
	current int
	start   int
}

func NewScanner(source string) *Scanner {
	return &Scanner{source: source}
}

func (s *Scanner) ScanTokens() []token.Token {
	for !s.atTheEnd() {
		tok := s.scanToken()
		s.tokens = append(s.tokens, tok)

		s.start = s.current
	}

    s.tokens = append(s.tokens, token.Token{token.EOF, "", ""})
	return s.tokens
}

func (s *Scanner) scanToken() token.Token {
	r := s.nextRune()

	switch r {
	case '(':
		return createToken(token.LEFT_PAREN, r)
	case ')':
		return createToken(token.RIGHT_PAREN, r)
	case '{':
		return createToken(token.LEFT_BRACE, r)
	case '}':
		return createToken(token.RIGHT_BRACE, r)
	case ',':
		return createToken(token.COMMA, r)
	case '.':
		return createToken(token.DOT, r)
	case '-':
		return createToken(token.MINUS, r)
	case '+':
		return createToken(token.PLUS, r)
	case ';':
		return createToken(token.SEMICOLON, r)
	case '/':
		return createToken(token.SLASH, r)
	case '*':
		return createToken(token.STAR, r)
	}
    return token.Token{}
}

func (s *Scanner) atTheEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) nextRune() byte {
	r := s.source[s.current]
	s.current++
	return r
}

func createToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Lexeme: string(char), Literal: string(char)}
}
