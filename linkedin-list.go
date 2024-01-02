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
	fmt.Println(addTwoNumbers(2, 3))
}

func linkedListSample(l *LinkedList) LinkedList {
	newNode := &Node{Data: data}
	newNode.Next = l.Head
	l.Head = newNode

	if (l.Tail == nil) {
		l.Tail = newNode
	}
}