package expr_test

import (
	"fmt"
	"github.com/stianwa/expr"
	"log"
)

func ExampleCalc() {
	f, err := expr.Calc("4 + 18/(9-6)")
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("answer: %f\n", f)
}

func ExampleRPN() {
	rpn, err := expr.RPN("4 + 18/(9-6)")
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("rpn: %s\n", rpn)
}
