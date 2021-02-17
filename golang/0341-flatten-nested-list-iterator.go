package main

func main() {
	//c = Constructor()
}

type NestedIterator struct {
	list []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	nestedIterator := &NestedIterator{}
	list := make([]int, 0)
	helper(nestedList, &list) //convert nestedList into []int
	nestedIterator.list = list
	return nestedIterator
}

func helper(nestedList []*NestedInteger, list *[]int) {
	for i := 0; i < len(nestedList); i++ {
		if nestedList[i].IsInteger() {
			*list = append(*list, nestedList[i].GetInteger()) //if is integer, append to list
		} else {
			helper(nestedList[i].GetList(), list) //if not integer, recur helper process
		}
	}
	return
}

func (this *NestedIterator) Next() int {
	res := 0
	if len(this.list) > 0 {
		res = this.list[0]
		this.list = this.list[1:]
	}
	return res
}

func (this *NestedIterator) HasNext() bool {
	if len(this.list) > 0 {
		return true
	}
	return false
}
