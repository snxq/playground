package queue

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	cases := []struct {
		Q      Queue
		Item   int
		Expect Queue
	}{
		{
			Q:      []int{},
			Item:   1,
			Expect: []int{1},
		}, {
			Q:      []int{1, 2},
			Item:   2,
			Expect: []int{1, 2, 2},
		},
	}

	for _, c := range cases {
		c.Q.Push(c.Item)
		if !reflect.DeepEqual(c.Q, c.Expect) {
			t.Errorf("Expect %v, Got %v", c.Expect, c.Q)
		}
	}
}

func TestPop(t *testing.T) {
	cases := []struct {
		Q       Queue
		Item    int
		Expect  Queue
		isError bool
	}{
		{
			Q:       []int{},
			Item:    -1,
			Expect:  []int{},
			isError: true,
		}, {
			Q:       []int{1},
			Item:    1,
			Expect:  []int{},
			isError: false,
		},
	}

	for _, c := range cases {
		actual, err := c.Q.Pop()
		if (err != nil) != c.isError {
			t.Errorf("Error should be %v, Got %v", c.isError, !c.isError)
		}
		if actual != c.Item {
			t.Errorf("Expect %d, Got %d", c.Expect, actual)
		}
	}
}
