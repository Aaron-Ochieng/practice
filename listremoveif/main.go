package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func insert(list *LinkedList, data int) {
	input := &Node{data, nil}

	if list.head == nil {
		list.head = input
		return
	}

	current := list.head
	for ; current.next != nil; current = current.next {
	}
	current.next = input
}

func printlist(list *LinkedList) {
	current := list.head

	for ; current != nil; current = current.next {
		fmt.Print(current.data)
		fmt.Print("--->")
	}
	fmt.Println(nil)
}

func remove(list *LinkedList, toberemoved int) {
	current := list.head
	for current.next != nil {
		if current.next.data == toberemoved {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func main() {
	ll := &LinkedList{}
	insert(ll, 200)
	insert(ll, 1)
	insert(ll, 2)
	insert(ll, 3)
	insert(ll, 10)
	insert(ll, 50)
	printlist(ll)
	remove(ll, 3)
	printlist(ll)
}
