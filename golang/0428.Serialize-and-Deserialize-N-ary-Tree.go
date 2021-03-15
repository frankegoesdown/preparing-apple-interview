package main

func main() {
	c := Constructor()
	a := c.deserialize()
}

type Node struct {
	Val      int
	Children []*Node
}

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

func (this *Codec) serialize(root *Node) string {
	stack := make([]string, 0)

	var helper func(*Node)
	helper = func(root *Node) {

		if root == nil {
			return
		} else {
			stack = append(stack, strconv.Itoa(root.Val))
			stack = append(stack, strconv.Itoa(len(root.Children)))

			for _, child := range root.Children {
				helper(child)
			}
		}
	}

	helper(root)
	return strings.Join(stack, ",")

}

func (this *Codec) deserialize(data string) *Node {

	if len(data) == 0 {
		return nil
	}
	dq := strings.Split(data, ",")

	var helper func() *Node
	helper = func() *Node {
		node := &Node{}
		value, _ := strconv.Atoi(dq[0])
		node.Val = value
		dq = dq[1:]

		size, _ := strconv.Atoi(dq[0])
		dq = dq[1:]
		for i := 0; i < size; i++ {
			node.Children = append(node.Children, helper())
		}
		return node
	}
	return helper()
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
