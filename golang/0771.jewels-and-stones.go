package main

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}

func numJewelsInStones(J string, S string) int {
	count := 0
	set := make(map[byte]bool)
	for i := 0; i < len(J); i++ {
		set[J[i]] = true
	}
	for i := 0; i < len(S); i++ {
		if _, ok := set[S[i]]; ok {
			count++
		}
	}
	return count
}
