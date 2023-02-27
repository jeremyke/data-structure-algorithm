package main

import "fmt"

func main() {
	fmt.Println(diGuiFindMax([]int{1, 6, 3, 24, 34, 7, 3}, 0, 6))
}

func diGuiFindMax(arr []int, l, r int) int {
	if l == r {
		return arr[l]
	}
	mid := (l + r) / 2
	lefMax := diGuiFindMax(arr, l, mid)
	rightMax := diGuiFindMax(arr, mid+1, r)
	return max(lefMax, rightMax)
}

func max(a, b int) int {
	res := a
	if b > a {
		res = b
	}
	return res
}
