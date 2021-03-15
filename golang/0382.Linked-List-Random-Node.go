package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := Constructor(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	})

	fmt.Println(c.GetRandom())
	fmt.Println("================")
	fmt.Println(c.GetRandom())
	fmt.Println("================")
	fmt.Println(c.GetRandom())
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	rand.Seed(time.Now().UnixNano())
	return Solution{
		head: head,
	}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	n := 1
	val := 0
	cur := this.head
	for cur != nil {
		fmt.Println(cur)
		fmt.Println(val)
		fmt.Println(n)
		fmt.Println("-----")
		if rand.Intn(n) == 0 {
			val = cur.Val
		}

		n++
		cur = cur.Next
	}

	return val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
