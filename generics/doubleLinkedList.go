package main

import (
	"errors"
	"fmt"
)

// Node Define the Node structure for doubly linked list
type Node[T any] struct {
	value      T
	next, prev *Node[T]
}

// DoublyLinkedList Define the DoublyLinkedList structure
type DoublyLinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

// Add appends the given value at the given index.
// Returns an error in the case that the index exceeds the list size.
func (l *DoublyLinkedList[T]) Add(index int, v T) error {
	if index > l.size {
		return errors.New("index exceeds list size")
	}
	newNode := &Node[T]{
		value: v,
	}
	if l.size == 0 {
		l.tail = newNode
		l.head = newNode
		l.size = 1
	} else {
		switch index {
		case 0:
			oldFirstNode := l.head
			oldFirstNode.prev = newNode
			l.head = newNode
			newNode.prev = nil
			newNode.next = oldFirstNode
		case l.size:
			oldLastNode := l.tail
			oldLastNode.next = newNode
			newNode.prev = oldLastNode
			newNode.next = nil
			l.tail = newNode
		default:
			currentNode := l.head
			for i := 0; i < index; i++ {
				currentNode = currentNode.next
				fmt.Printf("currentNode i: %v, val: %v\n", i+1, currentNode.value)
			}
			newNode.prev = currentNode.prev
			newNode.prev.next = newNode
			currentNode.prev = newNode
			newNode.next = currentNode
		}
		l.size += 1
	}
	fmt.Printf("in Add(%v, %v) :  list Forward : %s\n", index, v, l.PrintForward())
	return nil
}

func (l *DoublyLinkedList[T]) AddElements(elements []struct {
	index int
	value T
}) error {
	for _, e := range elements {
		fmt.Printf("In AddElements, size : %v about to Add index:%v, val: %v \n", l.size, e.index, e.value)
		if err := l.Add(e.index, e.value); err != nil {
			return err
		}
	}

	return nil
}
func (l *DoublyLinkedList[T]) PrintForward() string {
	if l.size == 0 {
		return ""
	}
	current := l.head
	// output := fmt.Sprintf("[%v] HEAD", l.head.value)
	output := "HEAD"
	for current != nil {
		output = fmt.Sprintf("%s -> %v", output, current.value)
		current = current.next
	}

	return fmt.Sprintf("%s -> NULL", output)
}

func (l *DoublyLinkedList[T]) PrintReverse() string {
	if l.size == 0 {
		return ""
	}
	current := l.tail
	output := "NULL"
	for current != nil {
		output = fmt.Sprintf("%s <- %v", output, current.value)
		current = current.prev
	}
	return fmt.Sprintf("%s <- HEAD", output)
}

func main() {

	testCases := []struct {
		index int
		value string
	}{
		{index: 0, value: "C"},
		{index: 0, value: "A"},
		{index: 1, value: "B"},
		{index: 3, value: "D"},
	}

	dl := &DoublyLinkedList[string]{}
	err := dl.AddElements(testCases)
	if err != nil {
		fmt.Println("got error doing AddElements :", err)
	}
	forwardPrint := dl.PrintForward()
	reversePrint := dl.PrintReverse()

	fmt.Println("END Forward : ", forwardPrint)
	fmt.Println("END Reverse : ", reversePrint)
}
