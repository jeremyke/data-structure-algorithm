package main

func mergeSort(arr []int, i, j int) []int {
	if i == j {
		return arr
	}
	mid := i + (j-i)/2
	mergeSort(arr, i, mid)
	mergeSort(arr, mid+1, j)
	merge(arr, i, mid, j)
	return arr
}

func merge(arr []int, i, mid, j int) {
	newArr := make([]int, 0, j-i+1) //定义一个数组能装i到j这么长的元素
	p1 := i
	p2 := mid + 1
	for p1 <= mid && p2 <= j {
		if arr[p1] > arr[p2] {
			newArr = append(newArr, arr[p2])
			p2++
		} else {
			newArr = append(newArr, arr[p1])
			p1++
		}
	}
	for p1 <= mid {
		newArr = append(newArr, arr[p1])
		p1++
	}
	for p2 <= mid {
		newArr = append(newArr, arr[p2])
		p2++
	}

	for m := 0; m < len(newArr); m++ {
		arr[i+m] = newArr[m]
	}
}
