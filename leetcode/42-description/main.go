package main

import "fmt"

/**
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/

func trap1(height []int) int {
	s := len(height)
	if s == 0 {
		return 0
	}
	var (
		leftHeightMax  = []int{height[0]}
		rightHeightMax = make([]int, s)
	)
	rightHeightMax[s-1] = height[s-1]
	for i := 1; i < s; i++ {
		leftHeightMax = append(leftHeightMax, max(leftHeightMax[i-1], height[i]))
	}
	for j := s - 2; j >= 0; j-- {
		rightHeightMax[j] = max(rightHeightMax[j+1], height[j])
	}
	res := 0
	for k, v := range height {
		res += min(leftHeightMax[k], rightHeightMax[k]) - v
	}
	return res
}

func trap(height []int) int {
	s := len(height)
	if s == 0 {
		return 0
	}
	var (
		res    = 0
		l      = 0
		r      = s - 1
		preMax = 0
		sufMax = 0
	)
	for l < r {
		preMax = max(preMax, height[l])
		sufMax = max(sufMax, height[r])
		if preMax < sufMax {
			res += preMax - height[l]
			l++
		} else {
			res += sufMax - height[r]
			r--
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
