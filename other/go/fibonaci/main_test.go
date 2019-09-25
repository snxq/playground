package fibonaci

import "testing"

func TestFibonaci(t *testing.T) {
	cases := []struct {
		Input  int
		Expect int
	}{
		{
			Input:  3,
			Expect: 2,
		}, {
			Input:  5,
			Expect: 5,
		},
	}

	for _, c := range cases {
		actual := Fibonaci(c.Input)
		if actual != c.Expect {
			t.Errorf("Expect %d, Got %d", c.Expect, actual)
		}
	}
}
