package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	var arr []int
	dfs(root, &arr)
	ans := &TreeNode{arr[0], nil, nil}
	tmp := ans
	for i := 1; i < len(arr); i++ {
		tmp.Right = &TreeNode{arr[i], nil, nil}
		tmp = tmp.Right
	}
	return ans
}

func dfs(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, arr)
	*arr = append(*arr, root.Val)
	dfs(root.Right, arr)
}
