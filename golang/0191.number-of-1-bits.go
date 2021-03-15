package main

func main() {

}

func hammingWeight(num uint32) int {
	var sum int
	for num != 0 {
		sum++
		num &= num - 1
	}
	return sum
}
