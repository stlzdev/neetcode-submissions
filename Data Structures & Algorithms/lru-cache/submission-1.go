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

func (this *LRUCache) remove(node *ListNode) {
	prev, nxt := node.Prev, node.Next
	prev.Next, nxt.Prev = nxt, prev
}

func (this *LRUCache) insert(node *ListNode) {
	prev, nxt := this.tail.Prev, this.tail
	prev.Next, nxt.Prev = node, node
	node.Next, node.Prev = nxt, prev
}

func (this *LRUCache) Get(key int) int {
    if node, exists := this.hmap[key]; exists {
		this.remove(this.hmap[key])
		this.insert(this.hmap[key])
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, exists := this.hmap[key]; exists {
		node.Val = value
		this.remove(this.hmap[key])
		this.insert(this.hmap[key])
		return
	}
	node := &ListNode{Key: key, Val: value}
	this.hmap[key] = node
	this.insert(node)
	if len(this.hmap) > this.capacity {
		lru := this.head.Next
		this.remove(lru)
		delete(this.hmap, lru.Key)
	}
}