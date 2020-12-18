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
		{Exp: "1 + 2 * 3 + 4 * 5 + 6", Sum: 71},
		{Exp: "1 + (2 * 3) + (4 * (5 + 6)) + 1", Sum: 52},
		{Exp: "2 * 3 + (4 * 5)", Sum: 26},
		{Exp: "5 + (8 * 3 + 9 + 3 * 4 * 3)", Sum: 437},
		{Exp: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", Sum: 12240},
		{Exp: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", Sum: 13632},
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
