package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) AddToEnd(value int) {
	newNode := &Node{Data: value}
	if list.Head == nil {
		list.Head = newNode
		return
	}

	currentNode := list.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = newNode
}

func (list *LinkedList) AddToHead(value int) {
	newNode := &Node{Data: value}
	if list.Head == nil {
		list.Head = newNode
		return
	}

	newNode.Next = list.Head
	list.Head = newNode
}

func (list *LinkedList) Print() {
	currentNode := list.Head
	for currentNode != nil {
		if currentNode.Next != nil {
			fmt.Printf("%v -> ", currentNode.Data)
			currentNode = currentNode.Next
		} else {
			fmt.Printf("%v\n", currentNode.Data)
			return
		}
	}
}

func (list *LinkedList) Size() int {
	size := 0
	currentNode := list.Head

	for currentNode != nil {
		size++
		currentNode = currentNode.Next
	}
	return size
}
