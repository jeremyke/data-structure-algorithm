package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr = []int{3, 7, 1, 5, 9, 4, 5}
	//fmt.Println(henan([]int{3, 7, 1, 5, 9, 4, 5}, 5))
	quickSortV3(arr, 0, 6)
	//fmt.Println(quickSortV2([]int{3, 5, 7, 1, 9, 5}, 0, 5))
	fmt.Println(arr)
}

func quickSortV3(arr []int, left, right int) []int {
	if left == right {
		return arr
	}
	roundIndex := getRandomWithAll(left, right)
	arr[roundIndex], arr[right] = arr[right], arr[roundIndex]
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
		quickSortV3(arr, left, smallRightIndex)
	}
	if bigLeftIndex <= right {
		quickSortV3(arr, bigLeftIndex, right)
	}

	return arr
}

func getRandomWithAll(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
