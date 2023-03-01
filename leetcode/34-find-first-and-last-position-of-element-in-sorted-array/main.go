package main

import "fmt"

/*
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
*/

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			start, stop := mid, mid
			for start > 0 {
				if target == nums[start-1] {
					start--
				} else {
					break
				}
			}
			for stop < len(nums)-1 {
				if target == nums[stop+1] {
					stop++
				} else {
					break
				}

			}
			return []int{start, stop}
		} else if target > nums[mid] {
			//tmpL := mid + 1
			//for nums[tmpL] == nums[mid+1] {
			//	tmpL++
			//}
			//l = tmpL - 1
			l = mid + 1

		} else {
			//tmpR := mid - 1
			//for nums[tmpR] == nums[mid-1] {
			//	tmpR--
			//}
			//r = tmpR + 1
			r = mid - 1
		}
	}
	return []int{-1, -1}
}

func main() {
	var nums = []int{5, 7, 7, 8, 8, 10}
	//var nums = []int{2, 2}
	fmt.Println(searchRange(nums, 8))
}
