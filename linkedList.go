package main

import "fmt"

type node struct {
	data int
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

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
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
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 16}
	node4 := &node{data: 7}
	node5 := &node{data: 9}
	node6 := &node{data: 2}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.printListData()
	mylist.deleteWithValue(100) //try delete value that does not exist
	mylist.deleteWithValue(2)   //Try delete the head Should work with out error
	mylist.printListData()
	emptyList := linkedList{}
	emptyList.deleteWithValue(10)

}
