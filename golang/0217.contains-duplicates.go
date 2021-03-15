package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}

func containsDuplicate(nums []int) bool {

	tmp := make(map[int]int)

	for _, v := range nums {

		tmp[v]++

		if tmp[v] > 1 {
			return true
		}
	}

	return false

}
