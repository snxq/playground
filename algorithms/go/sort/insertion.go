package sort

import (

)

func Insertion(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}
	var i, j int
	for i = 1; i < l; i++ {
		key := arr[i]
		for j = i-1; j >= 0; j-- {
			if key < arr[j] {
				arr[j+1] = arr[j]
				arr[j] = key
			} else {
				break
			}
		}
	}

	return arr
}