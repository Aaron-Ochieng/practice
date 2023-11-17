package linkedlists

import "fmt"

type Node struct {
	Next *Node
	Data interface{}
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Insert(data interface{}) {
	newNode := &Node{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func (l *LinkedList) PrintList() {
	for curent := l.Head; curent != nil; curent = curent.Next {
		fmt.Print(curent.Data, "-->")
	}
}

// func main() {
// 	mylist := LinkedList{}
// 	// mylist.insert("Hello World")
// 	// mylist.insert("Hello World")
// 	// mylist.insert("Hello World")
// 	// mylist.insert("Hello World")
// 	// mylist.insert("Hello World")
// 	mylist.printList()
// }
