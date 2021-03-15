package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	return splitLinkedList(copyNodeToLinkedList(head))
}

func splitLinkedList(head *Node) *Node {
	ptr := head
	head = head.Next
	for ptr != nil && ptr.Next != nil {
		ptr.Next, ptr = ptr.Next.Next, ptr.Next
	}
	return head
}

func copyNodeToLinkedList(head *Node) *Node {
	ptr := head
	for ptr != nil {
		node := &Node{
			Val:  ptr.Val,
			Next: ptr.Next,
		}
		ptr.Next, ptr = node, ptr.Next
	}
	ptr = head
	for ptr != nil {
		if ptr.Random != nil {
			ptr.Next.Random = ptr.Random.Next
		}
		ptr = ptr.Next.Next
	}
	return head
}
