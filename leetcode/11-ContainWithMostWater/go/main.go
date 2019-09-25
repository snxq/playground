package main

func maxArea(height []int) int {
	p1 := 0
	p2 := len(height) - 1
	maxArea := 0

	for {
		var area int
		if height[p1] > height[p2] {
			area = (p2 - p1) * height[p2]
			p2--
		} else {
			area = (p2 - p1) * height[p1]
			p1++
		}
		if area > maxArea {
			maxArea = area
		}
		if p1 == p2 {
			break
		}
	}

	return maxArea
}
