package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}
func sortArrayByParity(A []int) []int {
	l, r := 0, len(A)-1
	for l< r {
		if A[l] %2 > A[r] %2 {
			A[l], A[r] = A[r], A[l]
		}
		if A[l]%2==0 { l++}
		if A[r]%2==1 { r--}
		fmt.Println(l)
		fmt.Println(r)
		fmt.Println(A)
	}
	return A
}
