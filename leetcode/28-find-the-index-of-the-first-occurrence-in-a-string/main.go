package main

import "fmt"

/**
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。
*/
func strStr(haystack string, needle string) int {
	i, j, index := 0, 0, -1
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			if index == -1 {
				index = i
			}
			i++
			j++
		} else {
			index = -1
			if j > 0 {
				i = i - j + 1
				j = 0
			} else {
				i++
			}
		}
	}
	if index > -1 && j <= len(needle)-1 {
		return -1
	}
	return index
}

func main() {
	haystack, needle := "mississippi", "issip"
	fmt.Println(strStr(haystack, needle))
}
