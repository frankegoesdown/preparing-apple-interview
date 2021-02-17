package main

import "fmt"

func main() {
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4}))
}

func decompressRLElist(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i += 2 {
		tmp := make([]int, nums[i])
		for j := 0; j < nums[i]; j++ {
			tmp[j] = nums[i+1]
		}
		res = append(res, tmp...)
		fmt.Println(i)
		fmt.Println(tmp)
		fmt.Println(res)
	}
	return res
}
