package main

type StackOnLinkedList struct {
	List *LinkedList
	Size int
}

func NewEmptyStackOnLinkedList() *StackOnLinkedList {
	return &StackOnLinkedList{NewLinkedList(), 0}
}

func NewStackOnLinkedList(list *LinkedList) *StackOnLinkedList {
	return &StackOnLinkedList{list, list.Size()}
}

func (stack *StackOnLinkedList) Push(data int) {
	stack.List.AddToHead(data)
}

func (stack *StackOnLinkedList) Pop() (int, bool) {
	if stack.List.Head == nil {
		return 0, false
	}

	value := stack.List.Head.Data
	stack.List.Head = stack.List.Head.Next
	return value, true
}

func (stack *StackOnLinkedList) Peek() (int, bool) {
	if stack.List.Head == nil {
		return 0, false
	}
	return stack.List.Head.Data, true
}

func PrintStackOnLinkedList(stack *StackOnLinkedList) {
	stack.List.Print()
}
