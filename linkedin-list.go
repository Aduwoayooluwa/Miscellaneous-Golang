package main
import "fmt"

//  simple linked list in golang
type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func main() {
	list := &LinkedList{}
	linkedListSample(list, 1)
	
	for node:=list.Head; node != nil; node=node.Next {
		fmt.Println(node.Data)
	}
}

func linkedListSample(l *LinkedList, data int) {
	newNode := &Node{Data: data}
	newNode.Next = l.Head
	l.Head = newNode

	if (l.Tail == nil) {
		l.Tail = newNode
	}
}