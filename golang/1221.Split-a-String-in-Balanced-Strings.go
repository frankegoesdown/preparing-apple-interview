package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
}

func balancedStringSplit(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == s[0] {
			count++
		} else {
			count--
		}
		if count == 0 {
			res++
		}
	}
	return res
}
