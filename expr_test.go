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
		{expr: "13*12-16^(1/2)", expects: 152},
		{expr: "14-144^(1/2)-4", expects: -2},
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

func TestCalcFails(t *testing.T) {
	tests := []string{
		"1+2-",
		"*1+2",
		"3(4+2)",
		"(4+2)5",
		"(4+2)^",
		"(4+2)^6-",
		"(",
		")",
		"-4+++*8",
		"4**8",
		"4 8 -",
	}
	for i, test := range tests {
		f, err := Calc(test)
		if err == nil {
			t.Fatalf("test %d [%s] should have failed: result: %f", i, test, f)
		}
	}
}

func TestRPN(t *testing.T) {
	tests := []struct {
		expr    string
		expects string
	}{
		{expr: "4 + 18 / (9 - 6)", expects: "4 18 9 6 - / +"},
		{expr: "14-144^(1/2)-4", expects: "14 144 1 2 / ^ - 4 -"},
		{expr: "10+3*(2*(6+(10-5)))^(1/2)", expects: "10 3 2 6 10 5 - + * 1 2 / ^ * +"},
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
