package expr

import (
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		expr    string
		expects float64
	}{
		{expr: "4 + 18 / (9 -3)", expects: 7},
		{expr: "4 --4", expects: 8},
		{expr: "4 -+4", expects: 0},
		{expr: "-12", expects: -12},
		{expr: "12", expects: 12},
	}
	for i, test := range tests {
		f, err := Calc(test.expr)
		if err != nil {
			t.Fatalf("test %d [%s]: %v", i, test.expr, err)
		}
		if f != test.expects {
			t.Fatalf("test %d [%s]: expected %f got %f", i, test.expr, test.expects, f)
		}
	}
}

func TestRPN(t *testing.T) {
	tests := []struct {
		expr    string
		expects string
	}{
		{expr: "4 + 18 / (9 - 6)", expects: "4 18 9 6 - / +"},
	}
	for i, test := range tests {
		f, err := RPN(test.expr)
		if err != nil {
			t.Fatalf("test %d [%s]: %v", i, test.expr, err)
		}
		if f != test.expects {
			t.Fatalf("test %d [%s]: expected %q got %q", i, test.expr, test.expects, f)
		}
	}
}
