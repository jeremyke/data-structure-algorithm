package main

func searchMatrix(matrix [][]int, target int) bool {
	var arr []int
	if len(matrix) == 0 {
		return false
	}
	for _, v := range matrix {
		if v[0] <= target && v[len(v)-1] >= target {
			arr = v
		}
	}
	if arr == nil {
		return false
	}
	if len(arr) == 0 {
		return false
	}
	if len(arr) == 1 {
		if arr[0] == target {
			return true
		} else {
			return false
		}
	}
	l, r := 0, len(arr)
	for l <= r {
		if l == r {
			if arr[l] == target {
				return true
			} else {
				return false
			}
		}
		mid := (l + r) / 2
		if target == arr[mid] {
			return true
		} else if target > arr[mid] {
			if mid+1 < r {
				l = mid + 1
			} else {
				l = r
			}

		} else {
			if mid-1 > l {
				r = mid - 1
			} else {
				r = l
			}
		}
	}
	return false
}
