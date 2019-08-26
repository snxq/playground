package main

import "fmt"

func twoSum(nums []int, target int) []int {
	dict := map[int]int{}

	for i, num := range nums {
		j, ok := dict[target-num]
		if ok {
			return []int{j, i}
		} else {
			dict[num] = i
		}
	}
	return []int{}
}


func main() {
	result := twoSum([]int{1, 2, 3, 4}, 7)
	fmt.Println(result)
}
