package main

/*
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。
*/

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//var arr = []int{4, -7, 2, 8, 1, 6, 0, 19, 3}
	var arr = []int{-1, 2, 1, -4}
	res := threeSumClosest(arr, 1)
	fmt.Println(res)
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var res = math.MaxInt64
	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			total := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(res-target)) > math.Abs(float64(total-target)) {
				res = total
			}
			if total < target {
				l++
			} else if total > target {
				r--
			} else {
				return total
			}
		}

	}
	return res
}
