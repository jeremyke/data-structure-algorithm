package main

import (
	"fmt"
	"sort"
)

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。
*/
func main() {
	var arr = []int{4, -7, 2, 8, 1, 6, 0, 19, 3}
	res := threeSum(arr, 1)
	fmt.Println(res)
}

// threeSum 该方法运行超时
func threeSum(arr []int, target int) [][]int {
	//先排序
	sort.Ints(arr)
	var res [][]int
	for i := 0; i < len(arr)-2; i++ {
		if arr[i] > target || i > 0 && arr[i] == arr[i-1] {
			continue
		}
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[j] > target-arr[i] || i > 0 && arr[i] == arr[i-1] {
				continue
			}
			for k := j + 1; k < len(arr); k++ {
				if arr[k] > target-arr[i]-arr[j] {
					continue
				}
				if arr[i]+arr[k]+arr[j] == target {
					res = append(res, []int{arr[i], arr[j], arr[k]})
				}
			}
		}
	}
	fmt.Println(arr)
	return res
}

func threeSum2(nums []int) [][]int {
	//先排序
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 || i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[j] > 0-nums[i] || j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[k] > 0-nums[i]-nums[j] || k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				if nums[i]+nums[k]+nums[j] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	fmt.Println(nums)
	return res
}
