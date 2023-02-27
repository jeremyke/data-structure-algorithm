package main

import "fmt"

func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
	return arr
}
func main() {
	fmt.Println(insertSort([]int{3, 7, 1, 9, 4, 2}))
}
