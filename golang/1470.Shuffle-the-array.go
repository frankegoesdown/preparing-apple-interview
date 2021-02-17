package main

import "fmt"

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}

func shuffle(nums []int, n int) []int {
	var ans []int
	for i, j := 0, n; i < n && j < 2*n; i, j = i+1, j+1 {
		ans = append(ans, []int{nums[i], nums[j]}...)
	}
	return ans
}
