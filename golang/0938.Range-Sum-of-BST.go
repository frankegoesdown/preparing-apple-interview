package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {

	if root == nil {
		return 0
	}

	var ans int

	dfs(root, L, R, &ans)
	return ans
}

func dfs(root *TreeNode, L int, R int, ans *int) {
	if root == nil {
		return
	}
	if L <= root.Val && root.Val <= R {
		*ans += root.Val
	}
	if L < root.Val {
		dfs(root.Left, L, R, ans)
	}
	if R > root.Val {
		dfs(root.Right, L, R, ans)
	}
}
