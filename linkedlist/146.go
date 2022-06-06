package linkedlist

type LRUCache struct {
	ht    map[int]*Node
	cache *DList
	cap   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		ht:    make(map[int]*Node, 0),
		cache: New(),
		cap:   capacity,
	}
}

// 抽象API
// 将某个 key 提升为最近使用的
func (c *LRUCache) makeRecently(key int) {
	node := c.ht[key]
	// 先从链表中删除这个节点
	c.cache.remove(node)
	// 重新插到队尾
	c.cache.addLast(node)
}

// 添加最近使用的元素
func (c *LRUCache) addRecently(key, val int) {
	node := newNode(key, val)
	// 链表尾部就是最近使用的元素
	c.cache.addLast(node)
	// 在 ht 中添加 key 的映射
	c.ht[key] = node
}

// 删除某一个 key
func (c *LRUCache) deleteKey(key int) {
	node, ok := c.ht[key]
	if !ok {
		return
	}
	// 从链表中删除
	c.cache.remove(node)
	// 从 ht 中删除
	delete(c.ht, key)
}

// 删除最久未使用的元素
func (c *LRUCache) removeLeastRecently() {
	// 链表头部的第一个元素就是最久未使用的
	deleteNode := c.cache.removeFirst()
	// 从 ht 中删除
	key := deleteNode.key
	delete(c.ht, key)
}

func (c *LRUCache) Get(key int) int {
	node, ok := c.ht[key]
	if !ok {
		return -1
	}
	// 将该数据提升为最近使用的
	c.makeRecently(key)
	return node.val
}

func (c *LRUCache) Put(key int, val int) {
	// key 已存在
	if _, ok := c.ht[key]; ok {
		// 删除旧的数据
		c.deleteKey(key)
		// 新插入的数据为最近使用的数据
		c.addRecently(key, val)
		return
	}
	if c.cap == c.cache.len {
		// 若容量已满，删除最久未使用的元素
		c.removeLeastRecently()
	}
	c.addRecently(key, val)
}

// 双向链表
type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

func newNode(k, v int) *Node {
	return &Node{
		key: k,
		val: v,
	}
}

type DList struct {
	head *Node
	tail *Node
	len  int
}

func New() *DList {
	head := newNode(0, 0)
	tail := newNode(0, 0)
	head.next = tail
	tail.prev = head
	return &DList{
		head: head,
		tail: tail,
		len:  0,
	}
}

// 在链表尾部添加节点 x，时间 O(1)
func (d *DList) addLast(x *Node) {
	x.prev = d.tail.prev
	x.next = d.tail
	d.tail.prev.next = x
	d.tail.prev = x
	d.len++
}

// 删除链表中的 x 节点（x 一定存在）
// 由于是双链表且给的是目标 Node 节点，时间 O(1)
func (d *DList) remove(x *Node) {
	x.prev.next = x.next
	x.next.prev = x.prev
	d.len--
}

// 删除链表中第一个节点，并返回该节点，时间 O(1)
func (d *DList) removeFirst() *Node {
	if d.head.next == d.tail {
		return nil
	}
	first := d.head.next
	d.remove(first)
	return first
}

// 返回链表长度，时间 O(1)
func (d *DList) Len() int {
	return d.len
}
