package main

func main() {

}

type MovingAverage struct {
	windowSize int
	nums       []int
	sum        int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{windowSize: size, nums: []int{}, sum: 0}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.nums) < this.windowSize {
		this.sum += val
		this.nums = append(this.nums, val)
	} else if len(this.nums) == this.windowSize {
		var top int
		top, this.nums = this.nums[0], this.nums[1:]
		this.sum -= top
		this.nums = append(this.nums, val)
		this.sum += val
	} else {
		panic("we shouldn't have more numbers than the window size")
	}

	return float64(this.sum) / float64(len(this.nums))
}
