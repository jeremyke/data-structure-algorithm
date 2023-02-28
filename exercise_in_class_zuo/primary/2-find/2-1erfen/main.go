package main

import (
	"fmt"
)

var index int

func main() {
	arr := []int{-6, -4, -1, 2, 7, 9, 12, 68}
	//fmt.Println(erFen(arr, 0, 7, 68))
	fmt.Println(findMostLeftNum(arr, 0, 7, 3))
}

func erFen(arr []int, l, r, target int) int {
	if l == r {
		if arr[l] == target {
			return l
		} else {
			return -1
		}
	}
	mid := (l + r) / 2
	if arr[mid] == target {
		return mid
	} else if target > arr[mid] {
		return erFen(arr, mid+1, r, target)
	} else {
		return erFen(arr, l, mid-1, target)
	}
}

func findMostLeftNum(arr []int, l, r, target int) int {
	if l == r {
		if arr[l] >= target {
			return l
		} else {
			return l + 1
		}
	}
	mid := (l + r) / 2
	if arr[mid] == target {
		return mid
	} else if target > arr[mid] {
		return findMostLeftNum(arr, mid+1, r, target)
	} else {
		return findMostLeftNum(arr, l, mid-1, target)
	}
}
