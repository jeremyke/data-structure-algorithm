package main

import (
	"math/rand"
	"time"
)

func quickSort(arr []int, l, r int) []int {
	if l == r {
		return arr
	}
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(r-l+1) + l
	arr[random], arr[r] = arr[r], arr[random]
	smallRightIndex := l - 1
	bigLeftIndex := r + 1
	gapNum := arr[r]
	for i := l; i < bigLeftIndex; {
		if arr[i] < gapNum {
			arr[smallRightIndex+1], arr[i] = arr[i], arr[smallRightIndex+1]
			smallRightIndex++
			i++
		} else if arr[i] == gapNum {
			i++
		} else {
			arr[bigLeftIndex-1], arr[i] = arr[i], arr[bigLeftIndex-1]
			bigLeftIndex--
		}
	}
	if smallRightIndex >= l {
		quickSort(arr, l, smallRightIndex)
	}
	if bigLeftIndex <= r {
		quickSort(arr, bigLeftIndex, r)
	}
	return arr
}
