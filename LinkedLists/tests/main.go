package main

import linkedlists "practice/LinkedLists"

func main() {
	mylist := linkedlists.LinkedList{}
	mylist.Insert("1")
	mylist.Insert("2")
	mylist.Insert("3")
	mylist.Insert("4")
	mylist.Insert("5")
	mylist.PrintList()
}
