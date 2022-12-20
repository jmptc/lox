package scanner

import (
	"fmt"

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
		s.scanToken()
		s.start = s.current
	}

	s.tokens = append(s.tokens, token.Token{Type: token.EOF, Lexeme: "", Literal: ""})
	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.nextChar()

	switch c {
	case '(':
		s.addToken(token.LEFT_PAREN)
	case ')':
		s.addToken(token.RIGHT_PAREN)
	case '{':
		s.addToken(token.LEFT_BRACE)
	case '}':
		s.addToken(token.RIGHT_BRACE)
	case ',':
		s.addToken(token.COMMA)
	case '.':
		s.addToken(token.DOT)
	case '-':
		s.addToken(token.MINUS)
	case '+':
		s.addToken(token.PLUS)
	case ';':
		s.addToken(token.SEMICOLON)
	case '/':
		s.addToken(token.SLASH)
	case '*':
		s.addToken(token.STAR)
	case '!':
		if s.nextMatch('=') {
			s.addToken(token.BANG_EQUAL)
		} else {
			s.addToken(token.BANG)
		}
	case '=':
		if s.nextMatch('=') {
			s.addToken(token.EQUAL_EQUAL)
		} else {
			s.addToken(token.EQUAL)
		}
	case '<':
		if s.nextMatch('=') {
			s.addToken(token.LESS_EQUAL)
		} else {
			s.addToken(token.LESS)
		}
	case '>':
		if s.nextMatch('=') {
			s.addToken(token.GREATER_EQUAL)
		} else {
			s.addToken(token.GREATER)
		}
	}
}

func (s *Scanner) atTheEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) nextChar() byte {
	r := s.source[s.current]
	s.current++
	return r
}

func (s *Scanner) nextMatch(expectedChar byte) bool {
	if s.atTheEnd() {
		return false
	}

	if s.source[s.current] != expectedChar {
		return false
	}

	s.current += 1
	return true
}

func (s *Scanner) addToken(tokenType token.TokenType) {
	text := s.source[s.start:s.current]
	fmt.Printf("addToken: %s\n", text)

	tok := createToken(tokenType, text)
	s.tokens = append(s.tokens, tok)
}

func createToken(tokenType token.TokenType, text string) token.Token {
	return token.Token{Type: tokenType, Lexeme: text, Literal: text}
}
