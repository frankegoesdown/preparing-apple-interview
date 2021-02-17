package main

import "fmt"

func main() {
	fmt.Println(countNegatives([][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}))
}

func countNegatives(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	res := 0
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] < 0 {
				res++
			} else {
				break
			}
		}
	}
	return res
}
