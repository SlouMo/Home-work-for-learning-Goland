package main

type StackOnLinkedListWithTail struct {
	List *LinkedList
	Size int
}

func NewEmptyStackOnLinkedWithTailList() *StackOnLinkedListWithTail {
	return &StackOnLinkedListWithTail{NewLinkedList(), 0}
}

func NewStackOnLinkedWithTailList(list *LinkedList) *StackOnLinkedListWithTail {
	return &StackOnLinkedListWithTail{list, list.Size()}
}

func (stack *StackOnLinkedListWithTail) Push(data int) {
	stack.List.AddToEnd(data)
}

func (stack *StackOnLinkedListWithTail) Pop() (int, bool) {
	if stack.List.Head == nil {
		return 0, false
	}

	if stack.List.Head.Next == nil {
		value := stack.List.Head.Data
		stack.List.Head = nil
		return value, true
	}

	prev := stack.List.Head
	current := stack.List.Head.Next

	for current.Next != nil {
		prev = current
		current = current.Next
	}

	prev.Next = nil
	return current.Data, true
}

func (stack *StackOnLinkedListWithTail) Peek() (int, bool) {
	current := stack.List.Head

	if current == nil {
		return 0, false
	}

	for current.Next != nil {
		current = current.Next
	}

	return current.Data, true
}

func PrintStackOnLinkedWithTailList(stack *StackOnLinkedListWithTail) {
	stack.List.Print()
}
