package main

import "fmt"

//以下均基于大根堆排序

func main() {
	//arr, num := []int{12, 7, 4}, 5
	////heapInsert
	//arr = append(arr, num)
	//heapInsert(arr, 3)
	//fmt.Println(arr)
	////heapify
	//arr[0], arr[len(arr)-1] = arr[len(arr)-1], arr[0]
	//heapify(arr, 0, 3)
	//fmt.Println(arr)
	var arr = []int{1, 3, 7, 6, 4, 2, 8}
	heapSort(arr)
	fmt.Println(arr)
}

func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] { //和父节点比，比父节点大就和父节点交换位置
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}

func heapify(arr []int, index, heapSize int) {
	left := index*2 + 1
	for left < heapSize {
		//取出2个子节点中最大的节点位置
		biggerChild := left
		if left+1 < heapSize && arr[left+1] > arr[left] {
			biggerChild = left + 1
		}

		if arr[index] <= arr[biggerChild] { //当前节点的数比较大子节点数小，互换

		} else if arr[index] > arr[biggerChild] { //当前节点的数比较大子节点数大，无需互换
			biggerChild = index
		}
		if biggerChild == index {
			break
		}
		//交换位置
		arr[biggerChild], arr[index] = arr[index], arr[biggerChild]
		index = biggerChild
		left = index*2 + 1
	}
}

func heapSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	//把数组放入一个大根堆里面
	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}
	heapSize := len(arr)
	//0位置的数和最后位置数交换
	arr[0], arr[heapSize-1] = arr[heapSize-1], arr[0]
	heapSize--
	//再把0位置数heapSize
	for heapSize > 0 {
		heapify(arr, 0, heapSize)
		arr[0], arr[heapSize-1] = arr[heapSize-1], arr[0]
		heapSize--
	}
}
