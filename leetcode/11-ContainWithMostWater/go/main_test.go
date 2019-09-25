package main

import "testing"

func TestMaxArea(t *testing.T) {
	cases := []struct {
		Nums   []int
		Expect int
	}{
		{
			Nums:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			Expect: 49,
		},
	}

	for _, c := range cases {
		actual := maxArea(c.Nums)
		if actual != c.Expect {
			t.Errorf("Expect %d, Got %d", c.Expect, actual)
		}
	}
}
