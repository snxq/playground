package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func compareSlice(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestInsertion(t *testing.T) {
	cases := []struct {
		Input  []int
		Expect []int
	}{
		{
			Input:  []int{5, 2, 4, 6, 1, 3},
			Expect: []int{1, 2, 3, 4, 5, 6},
		}, {
			Input:  []int{99, 22, 13, 0, -1},
			Expect: []int{-1, 0, 13, 22, 99},
		}, {
			Input:  []int{3},
			Expect: []int{3},
		},
	}

	for _, c := range cases {
		actual := Insertion(c.Input)
		if !reflect.DeepEqual(actual, c.Expect) {
			t.Errorf("Expect %v, Got %v", c.Expect, actual)
		}
	}
}

func BenchmarkInsertion(b *testing.B) {
	rand.Seed(time.Now().Unix())
	arr := []int{}
	for i := 0; i < 80000; i++ {
		arr = append(arr, rand.Intn(100000))
	}
	for i := 0; i < b.N; i++ {
		Insertion(arr)
	}
}
