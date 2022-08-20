package main

import "fmt"

type node struct {
	data string
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++

}

func printNode(n *node) {
	fmt.Printf("%s ", n)
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%s ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value string) {
	//case linkList has no members
	if l.length == 0 {
		return
	}
	//case value to delete is the head item
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	//case input value does not exsist inside the list

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		//case input value does not exsist inside the list
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: "server3434843"}
	node2 := &node{data: "server3403822"}
	node3 := &node{data: "server6985855"}
	node4 := &node{data: "server5455252"}
	node5 := &node{data: "server9034900"}
	node6 := &node{data: "server3r67263"}
	node7 := &node{data: "server2390932"}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node6)
	mylist.prepend(node7)
	mylist.printListData()
	mylist.deleteWithValue("server1209122") //try delete value that does not exist
	mylist.deleteWithValue("server2390932") //Try delete the head Should work with out error
	mylist.deleteWithValue("server5455252") //Delete normal
	mylist.printListData()
	emptyList := linkedList{}
	emptyList.deleteWithValue("server5398433") //Try to delete from empty list
	printNode(node5)
	mylist.prepend(node4) //re-prepend server to list after it was deleted
	fmt.Printf("\n")
	mylist.printListData()
}
