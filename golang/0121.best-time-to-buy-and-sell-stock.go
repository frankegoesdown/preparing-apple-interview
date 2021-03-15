package main

func main() {

}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	result := 0
	temp := 0
	for i := 1; i < len(prices); i++ {
		next := temp + prices[i] - prices[i-1]
		result = max(result, next)
		if next >= 0 {
			temp = next
		} else {
			temp = 0
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
