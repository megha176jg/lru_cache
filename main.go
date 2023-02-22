package main

import "fmt"

type Node struct {
	prev  *Node
	key   int
	value int
	next  *Node
}

var tail Node
var head Node
var mp map[int]*Node
var capacity int

func main() {
	var option int
	mp = make(map[int]*Node)
	head = Node{
		prev:  nil,
		key:   0,
		value: 0,
		next:  nil,
	}
	tail = Node{
		prev:  nil,
		key:   0,
		value: 0,
		next:  nil,
	}
	head.next = &tail
	tail.prev = &head
	fmt.Println("enter the capacity of cache")
	fmt.Scanln(&capacity)
	fmt.Println("enter \n 1 to put \n 2 to get \n 3 to show the cache \n 0 to exit")
	fmt.Scanln(&option)
	for option > 0 {
		if option == 1 {
			var x int
			var y int
			fmt.Println("enter the key and value")
			fmt.Scanln(&x)
			fmt.Scanln(&y)
			put(x, y)
		}
		if option == 2 {
			var x int
			fmt.Println("Enter the key")
			fmt.Scanln(&x)
			fmt.Println(get(x))
		}

		if option == 3 {
			node := *(head.next)
			for node != tail {
				fmt.Println("key is %d and value is %d \n", node.key, node.value)
				node = *node.next
			}
		}
		fmt.Scanln(&option)
	}
}

func put(x int, y int) *Node {
	_, ok := mp[x]
	if !ok {
		if len(mp) >= capacity {
			tail.prev = tail.prev.prev
			tail.prev.prev.next = &tail
		}
		var newNode Node
		newNode.key = x
		newNode.value = y
		newNode.next = head.next
		newNode.prev = &head
		head.next = &newNode
		newNode.next.prev = &newNode
		mp[x] = &newNode

	} else {
		node := mp[x]
		(*node).prev.next = (*node).next
		(*node).next.prev = (*node).prev
		var newNode Node
		newNode.key = x
		newNode.value = y
		newNode.next = head.next
		newNode.prev = &head
		head.next = &newNode
		newNode.next.prev = &newNode
		mp[x] = &newNode
	}
	return &head
}
func get(x int) int {
	y, ok := mp[x]
	if ok {
		res := y.value
		y.prev.next = y.next
		y.next.prev = y.prev
		y.next = head.next
		head.next = y
		y.next.prev = y
		y.prev = &head
		return res
	} else {
		return -1
	}
}
