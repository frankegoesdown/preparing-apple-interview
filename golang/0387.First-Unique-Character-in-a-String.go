package main

func main() {

}

func firstUniqChar(s string) int {
	tmp := make([]int, 26)
	for _, v := range s {
		tmp[v-'a']++
	}
	for i, v := range s {
		if tmp[v-'a'] == 1 {
			return i
		}
	}
	return -1
}
