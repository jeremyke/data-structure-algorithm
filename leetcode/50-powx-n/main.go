package main

import "fmt"

/**
实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。
*/

func main() {
	var x float64 = 2.00000
	var n int = 3
	fmt.Println(myPow(x, n))
}

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return process(x, n)
	}
	return 1 / process(x, -n)
}

func process(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := process(x, n/2)
	if n%2 == 0 {
		return y * y
	} else {
		return y * y * x
	}
}
