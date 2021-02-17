package main

func main() {

}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	} else if root.Val < val {
		return searchBST(root.Right, val)
	} else if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return root
	}
}
