package main

import (
	"fmt"
	"math"
)

func main() {
	var a = []int{4, 7, 14, 12, 2, 19, 15}
	newArr := radixSort(a, 0, 6, 2)
	fmt.Println(newArr)
}

func radixSort(arr []int, l, r, digit int) []int { //digit为最大数的位数
	radix := 10                  //词频数
	bucket := make([]int, r-l+1) //bucket用来装排序之后的arr
	for d := 1; d <= digit; d++ {
		countArr := make([]int, 10) //词频数组

		//入桶，此时词频数组里面的元素代表该数字出现的次数
		for i := l; i <= r; i++ {
			j := getDigit(arr[i], d)
			countArr[j]++
		}
		//把词频数组里面的元素从代表该数字出现的次数-->改为小于等于该数字的元素次数
		for x := 1; x < radix; x++ {
			countArr[x] = countArr[x] + countArr[x-1]
		}

		//遍历arr,放入新的数组里面
		//1.比如countArr[4] = 8,说明d位上，数字小于等于4的数字一共有8个，从右到左便历bucket，找到一个d位上为4的数字，则他一定位于第7位（0为首位），
		//2.把该数字放在复刻数组index为7的位置上，同时词频数组中的词频数减一，则为countArr[4] = 7
		for m := r; m >= l; m-- {
			n := getDigit(arr[m], d)
			bucket[countArr[n]-1] = arr[m]
			countArr[n]--
		}

		//把bucket里面的元素复制到原数组
		copy(arr, bucket)
	}
	return arr
}

//
// getDigit
//  @Description: 获取某个数某一位置上的数字
//  @param num
//  @param d
//  @return int
//
func getDigit(num, d int) int {
	return num / int(math.Pow(float64(10), float64(d-1))) % 10
}
