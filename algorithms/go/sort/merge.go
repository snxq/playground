package sort

// MergeSort merge sort
func MergeSort(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}
	left := arr[:l/2]
	right := arr[l/2:]

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) []int {
	var res []int
	for i := 0; len(left) > 0 && len(right) > 0; i++ {
		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		res = append(res, left...)
	} else {
		res = append(res, right...)
	}
	return res
}
