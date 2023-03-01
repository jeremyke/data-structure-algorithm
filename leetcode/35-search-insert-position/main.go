package main

import "fmt"

/**
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。
*/
func searchInsert(nums []int, target int) int {
	if len(nums) == 1 {
		if target > nums[0] {
			return 1
		} else {
			return 0
		}
	}
	l, r := 0, len(nums)-1
	for l <= r {
		if l == r {
			if target > nums[r] {
				return l + 1
			} else {
				return l
			}
		}
		mid := (l + r) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			if mid+1 < len(nums) && mid+1 <= r {
				l = mid + 1
			} else {
				r = l
			}

		} else {
			if mid-1 >= 0 && mid-1 >= l {
				r = mid - 1
			} else {
				r = l
			}
		}
	}
	return 0
}

func main() {
	//var nums = []int{1, 3, 5, 6}
	var nums = []int{3, 5, 7, 9, 10}
	fmt.Println(searchInsert(nums, 8))
}
