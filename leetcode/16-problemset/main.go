package main

/*
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。
*/

import "fmt"

func main() {
	var arr = []int{4, -7, 2, 8, 1, 6, 0, 19, 3}
	res := threeSumClosest(arr, 1)
	fmt.Println(res)
}

func threeSumClosest(nums []int, target int) int {
	return 0
}
