package main

func main() {

}

func pivotIndex(nums []int) int {

	sum := 0
	leftSum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		if leftSum == sum-leftSum-nums[i] {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}
