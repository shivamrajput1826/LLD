package main

import "sync"

type Node[K comparable, V any] struct {
	key   K
	value V
	prev  *Node[K, V]
	next  *Node[K, V]
}

type LRUCache[K comparable, V any] struct {
	capacity int
	cache    map[K]*Node[K, V]
	head     *Node[K, V]
	tail     *Node[K, V]
	mu       sync.RWMutex
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	cache := &LRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]*Node[K, V]),
	}
	cache.head = &Node[K, V]{}
	cache.tail = &Node[K, V]{}
	cache.head.next = cache.tail
	cache.tail.next = cache.head
	return cache
}

func (c *LRUCache[K, V]) removeNode(node *Node[K, V]) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache[K, V]) addToHead(node *Node[K, V]) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache[K, V]) moveToHead(node *Node[K, V]) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *LRUCache[K, V]) removeTail() *Node[K, V] {
	node := c.tail.prev
	c.removeNode(node)
	return node
}

func (c *LRUCache[K, V]) Size() int {
	c.mu.RLock()
	defer c.mu.Unlock()
	return len(c.cache)
}

func (c *LRUCache[K, V]) Clear() {
	c.cache = make(map[K]*Node[K, V])
	c.head.next = c.tail
	c.tail.prev = c.head
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	node, ok := c.cache[key]
	if ok {
		c.moveToHead(node)
		return node.value, true
	}
	var zero V
	return zero, false

}

func (c *LRUCache[K, V]) Put(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if node, exist := c.cache[key]; exist {
		node.value = value
		c.moveToHead(node)
		return
	}
	newNode := &Node[K, V]{
		key:   key,
		value: value,
	}
	c.cache[key] = newNode
	c.addToHead(newNode)
	if len(c.cache) > c.capacity {
		tail := c.removeTail()
		delete(c.cache, tail.key)

	}
}
