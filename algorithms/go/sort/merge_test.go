package sort

import (
	"math/rand"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		l := rand.Intn(100)
		// Initial case
		arr := []int{}
		for j := 0; j < l; j++ {
			arr = append(arr, rand.Intn(1000000))
		}
		// Sort
		res := MergeSort(arr)
		// Check
		var key int
		for j := 0; j < l; j++ {
			if res[j] < key {
				t.Errorf("pre array: %v\nsorted array: %v\nkey: %d\nj: %d", arr, res, key, res[j])
			}
			key = res[j]
		}
		t.Logf("Test case %d pass", i)
	}
}
