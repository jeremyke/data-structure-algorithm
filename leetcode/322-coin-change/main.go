package main

import (
	"fmt"
	"math"
)

/**
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。
*/
func coinChange1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := math.MaxInt64
	for _, coin := range coins {
		sub := coinChange1(coins, amount-coin)
		if sub == -1 {
			continue
		}
		res = min(res, sub+1)
	}
	if res == math.MaxInt64 {
		return -1
	} else {
		return res
	}
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 数组大小为 amount + 1，初始值也为 amount + 1
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	// base case
	dp[0] = 0
	// 外层 for 循环在遍历所有状态的所有取值
	for i := 0; i < len(dp); i++ {
		// 内层 for 循环在求所有选择的最小值
		for _, coin := range coins {
			// 子问题无解，跳过
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	coins, amount := []int{1, 2, 5}, 11
	fmt.Println(coinChange(coins, amount))
}
