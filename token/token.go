package token

import (
	"fmt"
)

type TokenType string

const (
	// Single-character tokens
	LEFT_PAREN = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE = "{"
	RIGHT_BRACE = "}"
	COMMA = ","
	DOT = "."
	MINUS = "-"
	PLUS = "+"
	SEMICOLON = ";"
    SLASH = "/"
	STAR = "*"

	// One or two character tokens
	BANG = "!"
	BANG_EQUAL = "!="
	EQUAL = "="
	EQUAL_EQUAL = "=="
	GREATER = ">"
	GREATER_EQUAL = ">="
	LESS = "<"
	LESS_EQUAL ="<="

	// Literals
	IDENTIFIER = "IDENTIFIER"
	STRING = "STRING"
	NUMBER = "NUM"

	// Keywords
	AND = "AND"
	CLASS = "CLASS"
	ELSE = "ELSE"
	FALSE = "FALSE"
	FUN = "FUN"
	FOR = "FOR"
	IF = "IF"
	NIL = "NIL"
	OR = "OR"
	PRINT = "PRINT"
	RETURN = "RETURN"
	SUPER = "SUPER"
	THIS = "THIS"
	TRUE = "TRUE"
	VAR = "VAR"
	WHILE = "WHILE"

	EOF = "EOF"
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal string
}

func (t *Token) string() string {
	return fmt.Sprintf("Token: %d", t.Type)
}
