package main

import "fmt"

//倒序
func selectSort(sortArr []int) {
	for i := 0; i < len(sortArr); i++ {
		for j := i + 1; j < len(sortArr); j++ {
			if sortArr[i] < sortArr[j] {
				sortArr[i], sortArr[j] = sortArr[j], sortArr[i]
			}
		}
	}
}

func main() {
	aa := []int{7, 3, 8, 1, 9, 4, 10}
	selectSort(aa)
	fmt.Println(aa)
}
