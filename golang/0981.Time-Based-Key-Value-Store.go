package main

import "fmt"

func main() {
	kv := Constructor()
	kv.Set("foo", "bar", 1) // store the key "foo" and value "bar" along with timestamp = 1
	fmt.Println(kv.Get("foo", 1)) // bar
	fmt.Println(kv.Get("foo", 3)) // bar
	kv.Set("foo", "bar2", 4)
	fmt.Println(kv.Get("foo", 4)) // bar2
	fmt.Println(kv.Get("foo", 5)) // bar2

}

type TimeMap struct {
	store map[string][]Node
}

type Node struct {
	i   int
	val string
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	tMap := TimeMap{}
	tMap.store = make(map[string][]Node)
	return tMap
}

func (this *TimeMap) Set(key string, value string, timestamp int) {

	this.store[key] = append(this.store[key], Node{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	val, ok := this.store[key]
	if !ok {
		return ""
	}
	i := 0
	j := len(val) - 1
	//fmt.Println(val)
	mid := 0
	for i <= j {
		mid = i + (j-i)/2
		if val[mid].i == timestamp {
			return val[mid].val
		}
		if (mid == len(val)-1 || val[mid+1].i > timestamp) && val[mid].i <= timestamp {
			return val[mid].val
		}
		if val[mid].i > timestamp {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}

	return ""
}
