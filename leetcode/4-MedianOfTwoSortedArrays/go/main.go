package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var p1, p2 int
	l := len(nums1) + len(nums2)

	nums3 := []int{}
	for i := 0; i < (l/2)+1; i++ {
		if p1 >= len(nums1) {
			nums3 = append(nums3, nums2[p2])
			p2++
		} else if p2 >= len(nums2) {
			nums3 = append(nums3, nums1[p1])
			p1++
		} else if nums1[p1] <= nums2[p2] {
			nums3 = append(nums3, nums1[p1])
			p1++
		} else if nums1[p1] > nums2[p2] {
			nums3 = append(nums3, nums2[p2])
			p2++
		}
	}
	l3 := len(nums3)
	if l%2 == 1 {
		return float64(nums3[l3-1])
	}
	return float64(nums3[l3-2]+nums3[l3-1]) / float64(2)
}
