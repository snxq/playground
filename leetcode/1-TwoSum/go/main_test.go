package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	result := twoSum([]int{1, 2, 3, 4}, 7)
	if result[0] != 2 && result[1] != 3 {
		t.Errorf("result is not correct.")
	}
}
