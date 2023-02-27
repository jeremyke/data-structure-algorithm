package main

import "fmt"

func main() {
	fmt.Println(quickSortV1([]int{3, 7, 8, 1, 4, 9, 5}, 0, 6))
}

func quickSortV1(arr []int, left, right int) []int {
	if left == right {
		return arr
	}
	smallLeftIndex := left - 1
	midNum := arr[right]
	for i := left; i < right; i++ {
		if arr[i] <= midNum {
			arr[smallLeftIndex+1], arr[i] = arr[i], arr[smallLeftIndex+1]
			smallLeftIndex++
		}
	}
	arr[smallLeftIndex+1], arr[right] = arr[right], arr[smallLeftIndex+1]
	if smallLeftIndex >= left {
		quickSortV1(arr, left, smallLeftIndex)
	}
	if smallLeftIndex < left {
		quickSortV1(arr, smallLeftIndex+1, right)
	}

	return arr
}
