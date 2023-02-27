package main

import "fmt"

type testFun func(int32) int32

func A(a int32) int32 {
	return a + int32(1)
}

func editFun(afun testFun) {
	fmt.Printf("afun传进来的-value:%v,afun传进来的-地址:%v\n", afun, &afun)
	bfun := func(a int32) int32 {
		return a + 3
	}
	fmt.Printf("bfun-value:%v\n", bfun)
	afun = bfun
	fmt.Printf("afun传进来修改后的-value:%v,afun传进来修改后的-地址:%v\n", afun, &afun)
	return
}

func main() {
	var afun = A
	fmt.Printf("afun原来的-value:%v,afun原来的-地址:%v\n", afun, &afun)
	editFun(afun)
	fmt.Printf("afun修改后后外层的-value:%v,afun修改后后外层的-地址:%v\n", afun, &afun)

	res := afun(3)
	fmt.Println(res)
}

func main1() {
	if true {
		defer fmt.Printf("1")
	} else {
		defer fmt.Printf("2")
	}
	fmt.Printf("3")
}
