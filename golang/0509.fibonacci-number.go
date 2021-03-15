package main

import "fmt"

func main() {
	fmt.Println(fib(5))
}

func fib(n int) int {
	nums := make(map[int]int, n+1)
	//f(n) = f(n-1)+f(n-2)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n < 0 {
		return n
	}
	nums[0] = 0
	nums[1] = 1
	fibHelp(n, nums)
	return nums[n]
}

func fibHelp(n int, nums map[int]int) {
	//Use ok idiom to check if a value exists in the map
	v1, ok1 := nums[n-1]
	v2, ok2 := nums[n-2]
	if ok1 && ok2 {
		nums[n] = v1 + v2
	} else if ok2 {
		fibHelp(n-1, nums)
		nums[n] = nums[n-1] + v2
	} else {
		fibHelp(n-2, nums)
		fibHelp(n-1, nums)
		nums[n] = nums[n-2] + nums[n-1]
	}

}
