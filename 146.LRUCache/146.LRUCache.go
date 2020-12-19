package main

import "fmt"

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}
type LRUCache struct {
	cap  int
	m    map[int]*Node
	head *Node
	tail *Node
}

func (lru *LRUCache) nodeToHead(node *Node) {
	lru.head.next.prev = node
	(*node).next = lru.head.next
	lru.head.next = node
	(*node).prev = lru.head
}

func deleteNode(node *Node) {
	(*node).next.prev = (*node).prev
	(*node).prev.next = (*node).next
}

func Constructor(capacity int) LRUCache {
	ch := LRUCache{
		cap: capacity,
		m:   make(map[int]*Node),
		head: &Node{
			key:   0,
			value: 0,
		},
		tail: &Node{
			key:   0,
			value: 0,
		},
	}
	(*ch.head).next = ch.tail
	(*ch.tail).prev = ch.head
	return ch
}

func (this *LRUCache) Get(key int) int {
	if node, okay := this.m[key]; okay == true {
		deleteNode(node)
		this.nodeToHead(node)
		return (*node).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, okay := this.m[key]; okay == true {
		(*node).value = value
		deleteNode(node)
		this.nodeToHead(node)
	} else {
		newNode := &Node{
			key:   key,
			value: value,
		}
		if len(this.m) == this.cap {
			evictKey := (*this.tail.prev).key
			delete(this.m, evictKey)
			deleteNode(this.tail.prev)
		}
		this.m[key] = newNode
		this.nodeToHead(newNode)
	}
}
