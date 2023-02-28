package main

import (
	"fmt"
	"math"
)

//无序数组求最大值

func findMax(arr []int, l, r int) int {
	if l == r {
		return arr[l]
	}
	mid := (l + r) / 2
	leftMax := findMax(arr, l, mid)
	rightMax := findMax(arr, mid+1, r)
	return int(math.Max(float64(leftMax), float64(rightMax)))
}

func main() {
	arr := []int{-6, -4, -1, 2, 7, 9, 12, 68}
	fmt.Println(findMax(arr, 0, 7))
}
