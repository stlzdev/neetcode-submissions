type LRUCache struct {
    hmap map[int]*ListNode
	head *ListNode
	tail *ListNode
	capacity int
}

type ListNode struct {
	Key int
	Val int
	Prev *ListNode
	Next *ListNode
}

func Constructor(capacity int) LRUCache {
	head := &ListNode{}
	tail := &ListNode{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		hmap: map[int]*ListNode{},
		head: head,
		tail: tail,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
    if node, exists := this.hmap[key]; exists {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		node.Prev = this.head
		node.Next = this.head.Next
		this.head.Next.Prev = node
		this.head.Next = node
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, exists := this.hmap[key]; exists {
		node.Val = value
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		node.Prev = this.head
		node.Next = this.head.Next
		this.head.Next.Prev = node
		this.head.Next = node
		return
	}
	node := &ListNode{Key: key, Val: value}
	this.hmap[key] = node
	node.Prev = this.head
	node.Next = this.head.Next
	this.head.Next.Prev = node
	this.head.Next = node
	if len(this.hmap) > this.capacity {
		lru := this.tail.Prev
		lru.Prev.Next = lru.Next
		lru.Next.Prev = lru.Prev
		delete(this.hmap, lru.Key)
	}
}