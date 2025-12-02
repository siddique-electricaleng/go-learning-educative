package main

import (
	"fmt"
	"unsafe"
)

// Struct without a pointer
type T1 struct {
	a, b int64
}

// Linked List Node
type LinkedListNode struct {
	data    float64
	nextPtr *LinkedListNode
}

// Double Linked List Node
/*
type DoublyLinkedListNode struct {
	prPtr *DoublyLinkedListNode
	data  float64
	suPtr *DoublyLinkedListNode
}
*/
// Binary Tree Node
type BTNode struct {
	leftPtr  *BTNode
	data     float64
	rightPtr *BTNode
}

func advanceStructMain() {

	// Size of a struct
	var tmpLinkedList *LinkedListNode = new(LinkedListNode)
	fmt.Printf("Size of struct with a pointer:%+v\n", unsafe.Sizeof(*tmpLinkedList))

	var tmpAnotherLinkedList *T1 = new(T1)
	fmt.Printf("Size of struct with a pointer:%+v\n", unsafe.Sizeof(*tmpAnotherLinkedList))

}
