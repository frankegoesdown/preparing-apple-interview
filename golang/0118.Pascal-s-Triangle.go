package main

func main() {

}

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	res := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		var arr []int
		for j := 0; j < i+1; j++ {
			left, right := 0, 0
			if j-1 >= 0 {
				left = res[i-1][j-1]
			}
			if j < len(res[i-1]) {
				right = res[i-1][j]
			}
			arr = append(arr, left+right)
		}
		res = append(res, arr)
	}
	return res
}
