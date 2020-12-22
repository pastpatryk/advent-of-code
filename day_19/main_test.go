package main

import (
	"reflect"
	"regexp"
	"testing"
)

func TestParseRule(t *testing.T) {
	t.Run("parsing letter rule", func(t *testing.T) {
		text := "1: \"a\""

		id, rule, err := parseRule(text)
		assertNoError(t, err)
		assertID(t, id, 1)

		if rule.letter != "a" {
			t.Errorf("wanted rule letter: %q, got: %q", "a", rule.letter)
		}
	})

	t.Run("parsing single reference rule", func(t *testing.T) {
		text := "1: 2 3"

		id, rule, err := parseRule(text)
		assertNoError(t, err)
		assertID(t, id, 1)

		want := [][]int{[]int{2, 3}}
		if !reflect.DeepEqual(want, rule.references) {
			t.Errorf("got %v, wanted %v", rule.references, want)
		}
	})

	t.Run("parsing multiple reference rule", func(t *testing.T) {
		text := "1: 2 3 | 4 5"

		id, rule, err := parseRule(text)
		assertNoError(t, err)
		assertID(t, id, 1)

		want := [][]int{[]int{2, 3}, []int{4, 5}}
		if !reflect.DeepEqual(want, rule.references) {
			t.Errorf("got %v, wanted %v", rule.references, want)
		}
	})
}

func TestRuleStr(t *testing.T) {
	t.Run("single letter rule", func(t *testing.T) {
		grammar := Grammar(make(map[int]GrammarRule))
		grammar[0] = GrammarRule{letter: "a"}

		want := "a"
		got := grammar.ruleStr(0)

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("single nested rule", func(t *testing.T) {
		grammar := Grammar(make(map[int]GrammarRule))
		grammar[0] = GrammarRule{references: [][]int{[]int{1}}}
		grammar[1] = GrammarRule{letter: "a"}

		want := "(a)"
		got := grammar.ruleStr(0)

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("deep nested rule", func(t *testing.T) {
		grammar := Grammar(make(map[int]GrammarRule))
		grammar[0] = GrammarRule{references: [][]int{[]int{1, 2}}}
		grammar[1] = GrammarRule{references: [][]int{[]int{3, 3}, []int{4}}}
		grammar[2] = GrammarRule{letter: "a"}
		grammar[3] = GrammarRule{letter: "b"}
		grammar[4] = GrammarRule{letter: "c"}

		want := "((bb|c)a)"
		got := grammar.ruleStr(0)

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestGrammarRegexp(t *testing.T) {
	grammar := Grammar(make(map[int]GrammarRule))
	grammar[0] = GrammarRule{references: [][]int{[]int{1, 2}}}
	grammar[1] = GrammarRule{references: [][]int{[]int{3, 3}, []int{4}}}
	grammar[2] = GrammarRule{letter: "a"}
	grammar[3] = GrammarRule{letter: "b"}
	grammar[4] = GrammarRule{letter: "c"}

	r, err := grammar.Regexp()
	assertNoError(t, err)

	assertMatch(t, r, "bba")
	assertMatch(t, r, "ca")
}

func assertMatch(t *testing.T, r *regexp.Regexp, s string) {
	t.Helper()
	if !r.MatchString(s) {
		t.Errorf("should match %q, but doesnt", s)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func assertID(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("got id: %d, wanted: %d", got, want)
	}
}
