package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}

func singleNumber(nums []int) int {

	var n int

	for _, v := range nums {
		fmt.Println(n)
		fmt.Println(v)
		n = n ^ v
		fmt.Println(n)
		fmt.Println("============")
	}

	return n

}

func singleNumber2(nums []int) int {
	tmp := make(map[int]int)
	sum1 := 0
	sum2 := 0
	for _, v := range nums {
		sum1 += v
		tmp[v] = 1
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
	fmt.Println(tmp)
	for i := range tmp {
		sum2 += i
	}
	fmt.Println(sum2)
	return 2*sum2 - sum1
}
