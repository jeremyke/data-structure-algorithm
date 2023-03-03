package main

type LRUCache struct {
	Capacity int
	DataMap  map[int]int
	DataList *Node
}

type Node struct {
	key  int
	Pre  *Node
	Next *Node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		DataMap:  map[int]int{},
		DataList: nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.DataMap[key]; ok {
		return v
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.DataMap[key]; ok {
		this.DataMap[key] = value
	} else {
		if len(this.DataMap) == this.Capacity {

		}
		this.Capacity++
		this.DataMap[key] = value
	}

}

func (this *LRUCache) Top(key int) {
	if this.Capacity == 1 {
		return
	}
	head := this.DataList
	tail := this.DataList
	for tail.Next != nil {
		if tail.key == key {
			if tail.Next != nil {
				tail.Pre.Next = tail.Next
				tail.Next.Pre = tail.Pre
				break
			}
			tail = tail.Next
		}
	}
	tail.Next = head.Next
	this.DataList.Next = tail

}

//func (this *LRUCache) DelLast() {
//	head := this.DataList
//	tail := this.DataList
//	for tail.Next != nil {
//
//	}
//
//}
