package expr

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// tokenizer provides a tokeniser from input expression
type tokenizer struct {
	lexer *lexer
}

// next returns next token or error from lexer
func (t *tokenizer) next() (*token, error) {
	c := t.lexer.next()
	for c == ' ' {
		c = t.lexer.next()
	}

	if c == 0 {
		return nil, nil
	}
	if isOp(c) || isParen(c) {
		return &token{op: c}, nil
	}
	var buf []byte
	if unicode.IsDigit(rune(c)) || c == '.' {
		buf = append(buf, c)
	} else {
		return nil, fmt.Errorf("syntax error in position %d: expecting an operand or a digit, got '%c'", t.lexer.position, c)
	}
	for {
		c := t.lexer.next()
		if c == 0 {
			break
		}
		if !(unicode.IsDigit(rune(c)) || c == '.') {
			t.lexer.reverse()
			break
		}
		buf = append(buf, c)
	}

	value, err := strconv.ParseFloat(string(buf), 64)
	if err != nil {
		return nil, fmt.Errorf("bad number ending at position %d: %v", t.lexer.position, err)
	}

	return &token{value: value}, nil
}

// newTokenizer returns a tokenizer from expression input
func newTokenizer(expr string) *tokenizer {
	return &tokenizer{
		lexer: newLexer(expr),
	}
}

// token represents an expression token. It can be an operand, sign,
// parenthesis or a number.
type token struct {
	op    byte
	value float64
}

// opPreferences returns the tokens operand preference
func (t *token) opPreference() int {
	switch t.op {
	case '^':
		return 3
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}

	return -1
}

func (t *token) String() string {
	if t.op != 0 {
		return string(t.op)
	}

	return trimFloat(t.value)
}

func trimFloat(f float64) string {
	if f == 0 {
		return "0"
	}
	s := fmt.Sprintf("%f", f)
	if strings.Contains(s, ".") {
		s = strings.TrimRight(s, "0")
		s = strings.TrimSuffix(s, ".")
	}

	return s
}

func (t *token) isOp() bool {
	return isOp(t.op)
}

func (t *token) isNumber() bool {
	return t.op == 0
}

func (t *token) isSign() bool {
	return isSign(t.op)
}

func (t *token) isMinus() bool {
	return isMinus(t.op)
}

func (t *token) isLParen() bool {
	return isLParen(t.op)
}

func (t *token) isRParen() bool {
	return isRParen(t.op)
}

func (t *token) isParen() bool {
	return isParen(t.op)
}

func isMinus(b byte) bool {
	return b == '-'
}

func isPluss(b byte) bool {
	return b == '+'
}

func isSign(b byte) bool {
	return isPluss(b) || isMinus(b)
}

func isOp(b byte) bool {
	return b == '*' || b == '/' || b == '+' || b == '-' || b == '^'
}

func isLParen(b byte) bool {
	return b == '('
}

func isRParen(b byte) bool {
	return b == ')'
}

func isParen(b byte) bool {
	return isLParen(b) || isRParen(b)
}
