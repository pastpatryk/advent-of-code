package main

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	cases := []struct {
		Exp string
		Sum int
	}{
		{Exp: "1 + (2 * 3) + (4 * (5 + 6))", Sum: 51},
		{Exp: "2 * 3 + (4 * 5)", Sum: 46},
		{Exp: "5 + (8 * 3 + 9 + 3 * 4 * 3)", Sum: 1445},
		{Exp: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", Sum: 669060},
		{Exp: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", Sum: 23340},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%q = %d", test.Exp, test.Sum), func(t *testing.T) {
			exp, err := parseExpression(test.Exp)
			got := exp.Eval()

			assertNoError(t, err)

			if got != test.Sum {
				t.Errorf("got %d, wanted %d", got, test.Sum)
			}
		})
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
