package main

import "fmt"

func listIslands(m [][]int) int {
	if m == nil || m[0] == nil {
		return 0
	}
	res := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 1 {
				res++
				affect(m, i, j)
			}
		}
	}
	return res
}

func affect(m [][]int, i, j int) {
	if i < 0 || i >= len(m) || j < 0 || j >= len(m[i]) || m[i][j] != 1 {
		return
	}
	m[i][j] = 2
	affect(m, i+1, j)
	affect(m, i-1, j)
	affect(m, i, j+1)
	affect(m, i, j-1)
}

func main() {
	var islands = [][]int{
		{0, 1, 1, 0, 1, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 1},
		{0, 1, 1, 0, 1, 0, 0},
	}
	res := listIslands(islands)
	fmt.Println(res)

}
