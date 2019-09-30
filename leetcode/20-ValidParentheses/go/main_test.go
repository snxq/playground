package main

import "testing"

func TestIsValid(t *testing.T) {
	cases := []struct {
		Input  string
		Expect bool
	}{
		{
			Input:  "()",
			Expect: true,
		}, {
			Input:  "()[]{}",
			Expect: true,
		}, {
			Input:  "(]",
			Expect: false,
		}, {
			Input:  "([]}",
			Expect: false,
		}, {
			Input:  "{[]}",
			Expect: true,
		}, {
			Input:  "",
			Expect: true,
		}, {
			Input:  "abc",
			Expect: false,
		}, {
			Input: "[",
			Expect: false,
		},
	}

	for _, c := range cases {
		actual := isValid(c.Input)
		if actual != c.Expect {
			t.Errorf("Expect %v, Got %v", c.Expect, actual)
		}
	}
}
