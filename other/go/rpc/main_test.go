package main

import "testing"

func TestMutiply(t *testing.T) {
	cases := []struct {
		A *Args
		R *int
	}{
		{
			A: &Args{A: 1, B: 2},
			R: new(int),
		}, {
			A: &Args{A: 0, B: 4},
			R: new(int),
		},
	}

	var a Arith
	for _, c := range cases {
		err := a.Mutiply(c.A, c.R)
		if err != nil {
			t.Error("Error not nil")
		}
	}
}
