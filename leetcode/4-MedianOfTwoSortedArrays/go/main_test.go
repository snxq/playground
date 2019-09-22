package main

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	testcases := []struct{
		Nums1 []int
		Nums2 []int
		Expect float64
	}{
		{
			Nums1: []int{1, 3},
			Nums2: []int{2},
			Expect: 2,
		},{
			Nums1: []int{1, 2},
			Nums2: []int{3, 4},
			Expect: 2.5,
		},{
			Nums1: []int{3, 4},
			Nums2: []int{1, 2},
			Expect: 2.5,
		},
	}

	for _, c := range testcases {
		result := findMedianSortedArrays(c.Nums1, c.Nums2)
		print(c.Expect)
		if result != c.Expect {
			t.Fatalf("Expect %f, Got %f", c.Expect, result)
		}
	}
}