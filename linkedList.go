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
	fmt.Printf("%s ", n.data)
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
	fmt.Println("Trying to Remove " + value + " from Server list")
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
	serveUp := linkedList{}
	serverDown := linkedList{}
	node1 := &node{data: "server3434843"}
	node2 := &node{data: "server3403822"}
	node3 := &node{data: "server6985855"}
	node4 := &node{data: "server5455252"}
	node5 := &node{data: "server9034900"}
	node6 := &node{data: "server3r67263"}
	node7 := &node{data: "server2390932"}
	serveUp.prepend(node1)
	serveUp.prepend(node2)
	serveUp.prepend(node3)
	serveUp.prepend(node4)
	serveUp.prepend(node6)
	serveUp.prepend(node7)
	serveUp.printListData()
	serveUp.deleteWithValue("server1209122") //try delete value that does not exist
	serveUp.deleteWithValue("server2390932") //Try delete the head Should work with out error
	serveUp.deleteWithValue("server5455252") //Delete normal
	serverDown.prepend(node4)
	serverDown.prepend(node7)
	fmt.Println("Server Up...")
	serveUp.printListData()
	fmt.Println("Server Down ...")
	serverDown.printListData()
	emptyList := linkedList{}
	emptyList.deleteWithValue("server5398433") //Try to delete from empty list
	printNode(node5)
	serveUp.prepend(node4) //re-prepend server to list after it was deleted
	serverDown.deleteWithValue("server2390932")
	fmt.Printf("\n")
	fmt.Println("Server Up...")
	serveUp.printListData()
	fmt.Println("Server Down...")
	serverDown.printListData()
}
