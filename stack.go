package expr

import (
	"fmt"
	"math"
)

type stack struct {
	tokens []*token
}

func (s *stack) push(t *token) {
	s.tokens = append(s.tokens, t)
}

func (s *stack) pop() *token {
	if len(s.tokens) == 0 {
		return nil
	}
	t := s.tokens[len(s.tokens)-1]
	s.tokens = s.tokens[0 : len(s.tokens)-1]
	return t
}

func (s *stack) top() *token {
	if len(s.tokens) == 0 {
		return nil
	}
	return s.tokens[len(s.tokens)-1]
}

func (s *stack) String() string {
	return fmt.Sprintf("%s", s.tokens)
}

func (s *stack) operation(op byte) error {
	v2 := s.pop()
	if v2 == nil {
		return fmt.Errorf("pop: stack empty")
	}
	v1 := s.pop()
	if v1 == nil {
		return fmt.Errorf("pop: stack empty")
	}
	var c float64
	switch op {
	case '^':
		c = math.Pow(v1.value, v2.value)
	case '*':
		c = v1.value * v2.value
	case '/':
		if v2.value == 0 {
			return fmt.Errorf("division by zero")
		}
		c = v1.value / v2.value
	case '+':
		c = v1.value + v2.value
	case '-':
		c = v1.value - v2.value
	default:
		return fmt.Errorf("unknown operator: '%c'", op)
	}
	s.push(&token{value: c})

	return nil
}
