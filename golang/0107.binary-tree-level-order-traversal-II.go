package main

import "fmt"

func main() {
	fmt.Println(levelOrderBottom(&TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	ans := [][]int{}
	for len(queue) > 0 {
		n := len(queue)
		tmp := []int{}
		for i := 0; i < n; i++ {
			pop := queue[0]
			queue = queue[1:]
			tmp = append(tmp, pop.Val)
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
		ans = append(ans, tmp)
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}
