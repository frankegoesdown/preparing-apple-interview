package main

func main() {

}

func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	return constructBST(nums, 0, len(nums)-1)

}

func constructBST(nums []int, left int, right int) *TreeNode {

	if left > right {
		return nil
	}

	mid := left + (right-left)/2

	return &TreeNode{
		Val:   nums[mid],
		Left:  constructBST(nums, left, mid-1),
		Right: constructBST(nums, mid+1, right),
	}

}
