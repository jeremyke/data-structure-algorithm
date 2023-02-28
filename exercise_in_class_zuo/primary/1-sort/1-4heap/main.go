package main

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
	swap(arr, 0, heapSize-1)
	//再把0位置数heapSize
	for heapSize > 0 {
		heapify(arr, 0, heapSize)
		swap(arr, 0, heapSize-1)
	}
}

func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		swap(arr, index, (index-1)/2)
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
		//把当前节点的数和子节点较大的数比较
		if arr[biggerChild] < arr[index] {
			biggerChild = index
		}
		if biggerChild == index {
			break
		}
		//交换位置
		swap(arr, biggerChild, index)
		index = biggerChild
		left = index*2 + 1
	}
}

func swap(arr []int, a, b int) {
	arr[a] = arr[a] ^ arr[b]
	arr[b] = arr[a] ^ arr[b]
	arr[a] = arr[a] ^ arr[b]
}
