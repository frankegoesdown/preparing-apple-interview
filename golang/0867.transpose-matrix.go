package main

import "fmt"

func main() {
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func transpose(A [][]int) [][]int {

	var tpMatrix [][]int

	for i := 0; i < len(A[0]); i++ {
		var tmp []int
		for I := 0; I < len(A); I++ {
			tmp = append(tmp, A[I][i])
		}
		tpMatrix = append(tpMatrix, tmp)
		fmt.Println(tpMatrix)
		fmt.Println(tmp)
	}

	return tpMatrix

}
