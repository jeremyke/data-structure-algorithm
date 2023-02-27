package main

import "fmt"

func main() {
	fmt.Println(erFen([]int{1, 2, 4, 7, 9, 13}, 0, 5, 13))
}

func erFen(arr []int, l, r, x int) int {
	if l == r {
		if arr[l] == x {
			return l
		} else {
			return -1
		}
	}
	mid := (l + r) / 2
	if arr[mid] == x {
		return mid
	} else if x > arr[mid] {
		return erFen(arr, mid+1, r, x)
	} else {
		return erFen(arr, l, mid-1, x)
	}
}
