package main

import "fmt"

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		fmt.Println(nums[nums[i]-1])
		fmt.Println(nums[i])
		for nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			fmt.Println(nums)
		}
		fmt.Println("=======")
	}
	fmt.Println(nums)
	var ret []int
	for i, v := range nums {
		fmt.Println(v)
		fmt.Println(i)
		if v != i+1 {
			ret = append(ret, i+1)
		}
		fmt.Println(ret)
	}

	return ret
}
