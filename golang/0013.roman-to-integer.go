package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

var valueMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) (sum int) {
	var greatest int

	for i := len(s) - 1; i >= 0; i-- {
		val := valueMap[s[i]]
		fmt.Println(val)
		fmt.Println(greatest)
		fmt.Println(sum)
		fmt.Println("========")
		if val >= greatest {
			greatest = val
			sum += val
			continue
		}
		sum -= val
	}
	return sum
}
