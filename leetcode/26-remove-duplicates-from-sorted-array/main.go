package main

import "fmt"

func removeDuplicates(nums []int) int {
	s := len(nums)
	if s == 0 {
		return 0
	} else if s == 1 {
		return 1
	}
	slow := 1
	for fast := 1; fast < s; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main() {
	var arr = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr))
}
