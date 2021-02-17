package main

import "fmt"

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}
func shortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)
	pos := -n
	for i := range ans {
		ans[i] = n
	}
	fmt.Println(n)
	fmt.Println(pos)
	fmt.Println(ans)
	fmt.Println("========")
	for i := 0; i < n; i++ {
		if s[i] == c {
			pos = i
		}
		fmt.Println(i)
		fmt.Println(pos)
		fmt.Println(ans)
		fmt.Println("----------")
		ans[i] = i - pos
	}
	fmt.Println(ans)
	fmt.Println("========")

	for i := pos - 1; i >= 0; i-- {
		if s[i] == c {
			pos = i
		}
		ans[i] = min(ans[i], pos-i)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
