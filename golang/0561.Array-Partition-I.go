package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrayPairSum([]int{6, 2, 6, 5, 1, 2}))
}

func arrayPairSum(nums []int) int {
	var sum int
	//sort slice
	sort.Ints(nums)

	//add the sum.  range is more efficient because we don't know the length
	for i, _ := range nums {
		if i%2 == 0 {
			sum += nums[i]
		}

	}

	return sum
}
