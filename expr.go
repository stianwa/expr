// Package expr implements parsing of basic infix arithmetic expressions.
package expr

import (
	"fmt"
	"strings"
)

// Calc parses and calculates the infix expression.
func Calc(expr string) (float64, error) {
	tokens, err := parse(expr)
	if err != nil {
		return 0, err
	}

	tokenRPN, err := shuntingYardAlgorithm(tokens)
	if err != nil {
		return 0, err
	}

	return calcRPN(tokenRPN)
}

// RPN converts an infix expression to a Reverse Polish Notation
// expression.
func RPN(expr string) (string, error) {
	tokens, err := parse(expr)
	if err != nil {
		return "", err
	}

	tokenRPN, err := shuntingYardAlgorithm(tokens)
	if err != nil {
		return "", err
	}

	var ret []string
	for _, token := range tokenRPN {
		ret = append(ret, token.String())
	}

	return strings.Join(ret, " "), nil
}

// parse parses expression into an infix token slice.
func parse(expr string) ([]*token, error) {
	tz := newTokenizer(expr)
	var tokens []*token

	negNext := false
	expectNumber := true
	balance := 0
	for {
		token, err := tz.next()
		if err != nil {
			return nil, err
		}
		if token == nil {
			break
		}
		if token.isLParen() {
			balance++
		} else if token.isRParen() {
			balance--
		} else if expectNumber {
			if token.isSign() {
				if token.isMinus() {
					negNext = true
				}
				continue
			}
			if !token.isNumber() {
				return nil, fmt.Errorf("expected number, got %q", token)
			}
			if negNext {
				token.value = token.value * -1
				negNext = false
			}
			expectNumber = !expectNumber
		} else {
			if !token.isOp() {
				return nil, fmt.Errorf("expected operand, got %q", token)
			}
			expectNumber = !expectNumber
		}

		tokens = append(tokens, token)
	}
	if balance != 0 {
		return nil, fmt.Errorf("unbalanced parenthesis")
	}
	if len(tokens) > 0 && tokens[len(tokens)-1].isOp() {
		return nil, fmt.Errorf("infix expression ending with operand")
	}

	return tokens, nil
}

// shuntingYardAlgorithm creates a reverse polish notation token slice
// from an infix slice
func shuntingYardAlgorithm(tokens []*token) ([]*token, error) {
	output, op := &stack{}, &stack{}
	for _, token := range tokens {
		if token.isOp() {
			pref := token.opPreference()
			top := op.top()
			if top == nil || pref >= top.opPreference() {
				op.push(token)
			} else if top != nil && pref < top.opPreference() {
				for t := op.pop(); t != nil; t = op.pop() {
					output.push(t)
				}
				op.push(token)
			} else {
				output.push(token)
			}
		} else if token.isLParen() {
			op.push(token)
		} else if token.isRParen() {
			for t := op.top(); !t.isLParen(); t = op.top() {
				output.push(op.pop())
			}
			op.pop()
		} else {
			// is number
			output.push(token)
		}
	}
	for token := op.pop(); token != nil; token = op.pop() {
		output.push(token)
	}

	return output.tokens, nil
}

// calcRPN executes calculation of RPN ordered tokens
func calcRPN(tokens []*token) (float64, error) {
	calc := &stack{}
	for _, token := range tokens {
		if token.isNumber() {
			calc.push(token)
			continue
		}
		if err := calc.operation(token.op); err != nil {
			return 0, err
		}
	}

	if len(calc.tokens) != 1 {
		return 0, fmt.Errorf("length of calc stack not 1: %s", calc)
	}

	return calc.top().value, nil
}
