package main

import "fmt"

func bubbleSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for e := len(arr) - 1; e > 0; e-- {
		for i := 0; i < e; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	return arr
}
func main() {
	fmt.Println(bubbleSort([]int{3, 7, 1, 9, 4, 2}))
}
