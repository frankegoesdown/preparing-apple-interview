package main

func main() {

}

func titleToNumber(s string) int {
	n, res := 1, 0
	for i := len(s) - 1; i >= 0; i-- {
		tmp := int(s[i]-'A') + 1
		res += tmp * n
		n *= 26
	}
	return res
}
