package main

import "fmt"

/**
整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。


*/
func nextPermutation(nums []int) {
	flag := false
	for j := len(nums) - 1; j >= 0; j-- {
		for i := j - 1; i >= 0; {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				i = j - 1
				flag = true
			} else {
				i--
			}
		}
		if flag {
			break
		}
	}
	if !flag {
		for m := 0; m < (len(nums)-1)/2; m++ {
			nums[m], nums[len(nums)-1-m] = nums[len(nums)-1-m], nums[m]
		}
	}
}

func main() {
	var arr = []int{1, 3, 2}
	nextPermutation(arr)
	fmt.Println(arr)
}
