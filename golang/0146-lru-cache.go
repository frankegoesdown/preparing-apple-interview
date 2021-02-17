package main

import (
	"container/list"
	"fmt"
)

func main() {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println(c.Get(1))
	c.Put(3, 3)
	fmt.Println(c.Get(2))

	c.Put(4, 4)

	fmt.Println(c.Get(1))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(4))
}

type LRUCache struct {
	capacity int
	list     *list.List
	cache    map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[int]*list.Element, capacity),
	}
}

func (l *LRUCache) Get(key int) int {
	element, has := l.cache[key]
	if !has {
		return -1
	}
	l.list.MoveBefore(element, l.list.Front())
	return element.Value.([]int)[1]
}

func (l *LRUCache) Put(key int, value int) {
	element, has := l.cache[key]
	if has {
		element.Value = []int{key, value}
		l.list.MoveBefore(element, l.list.Front())
		return
	}
	if l.list.Len()+1 > l.capacity {
		back := l.list.Back()
		l.list.Remove(back)
		delete(l.cache, back.Value.([]int)[0])
	}
	front := l.list.PushFront([]int{key, value})
	l.cache[key] = front
}
