package main

import "fmt"

func main() {
	//fmt.Println(henan([]int{3, 7, 1, 5, 9, 4, 5}, 5))
	fmt.Println(quickSortV2([]int{3, 7, 1, 5, 9, 4, 5}, 0, 6))
	//fmt.Println(quickSortV2([]int{3, 5, 7, 1, 9, 5}, 0, 5))
}

func quickSortV2(arr []int, left, right int) []int {
	if left == right {
		return arr
	}
	smallRightIndex := left - 1
	bigLeftIndex := right + 1
	gapNum := arr[right]
	for i := left; i < bigLeftIndex; {
		if arr[i] < gapNum {
			arr[smallRightIndex+1], arr[i] = arr[i], arr[smallRightIndex+1]
			smallRightIndex++
			i++
		} else if arr[i] == gapNum {
			i++
		} else if arr[i] > gapNum {
			arr[bigLeftIndex-1], arr[i] = arr[i], arr[bigLeftIndex-1]
			bigLeftIndex--
		}

	}
	if smallRightIndex >= left {
		quickSortV2(arr, left, smallRightIndex)
	}
	if bigLeftIndex <= right {
		quickSortV2(arr, bigLeftIndex, right)
	}

	return arr
}

func henan(arr []int, num int) []int {

	smallRightIndex := -1
	bigLeftIndex := len(arr)

	for i := 0; i < bigLeftIndex; {
		if arr[i] < num {
			arr[smallRightIndex+1], arr[i] = arr[i], arr[smallRightIndex+1]
			smallRightIndex++
			i++
		} else if arr[i] == num {
			i++
		} else if arr[i] > num {
			arr[bigLeftIndex-1], arr[i] = arr[i], arr[bigLeftIndex-1]
			bigLeftIndex--
		}

	}
	return arr
}
