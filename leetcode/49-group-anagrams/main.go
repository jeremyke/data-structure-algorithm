package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	strMap := map[string][]string{}
	for _, v := range strs {
		vStrArr := []byte(v)
		sort.Slice(vStrArr, func(i, j int) bool {
			return vStrArr[i] < vStrArr[j]
		})
		vStr := string(vStrArr)
		strMap[vStr] = append(strMap[vStr], v)
	}
	res := make([][]string, 0, len(strMap))
	for _, v := range strMap {
		res = append(res, v)
	}
	return res
}

func main() {
	var arr = []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	res := groupAnagrams(arr)

	fmt.Println(res)

}
