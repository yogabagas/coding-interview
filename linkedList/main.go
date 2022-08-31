package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (l *LinkedList) Display() {

	for l.head != nil {
		fmt.Println("-->", l.head.value)
		l.head = l.head.next
	}
	fmt.Println()

}

func constructNewNode(value int) *Node {
	return &Node{
		value: value,
		next:  nil,
	}
}

func (l *LinkedList) Append(value int) {
	newNode := constructNewNode(value)

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.length++
	} else {
		l.tail.next = newNode
		l.tail = newNode
		l.length++
	}
}

func (l *LinkedList) Prepend(value int) {
	newNode := constructNewNode(value)
	newNode.next = l.head
	l.head = newNode
	l.length++
}

func (l *LinkedList) Insert(indexAt int, value int) {
	newNode := constructNewNode(value)

	if indexAt >= l.length {
		l.Append(value)
	}

	leader := l.TraverseToIndex(indexAt - 1)
	holdingPointer := leader.next
	leader.next = newNode
	newNode.next = holdingPointer
	l.length++
}

func (l *LinkedList) Delete(index int) {

	leader := l.TraverseToIndex(index - 1)
	unwantedPointer := leader.next
	leader.next = unwantedPointer.next

	l.length--

}

func (l *LinkedList) TraverseToIndex(index int) *Node {
	var counter int
	var currentNode = l.head

	for counter != index {
		currentNode = currentNode.next
		counter++
	}
	return currentNode
}

func (l *LinkedList) Length() int {
	return l.length
}

func main() {
	var l LinkedList

	l.Append(5)

	l.Append(6)

	l.Append(10)

	l.Prepend(1)

	l.Insert(2, 9)

	l.Delete(1)

	fmt.Println("head:", l.head.value)
	fmt.Println("tail:", l.tail.value)
	fmt.Println("length:", l.length)

	l.Display()
}
