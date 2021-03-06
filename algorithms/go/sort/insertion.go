package sort

// Insertion insert sort
func Insertion(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}

	var i, j int
	for i = 1; i < l; i++ {
		key := arr[i]
		for j = i - 1; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}

	return arr
}
