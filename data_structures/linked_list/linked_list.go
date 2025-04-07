package linked_list

import "fmt"

type Node[T comparable] struct {
	next  *Node[T]
	value T
}

type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
}

func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{
		value: value,
	}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.next = newNode
		ll.tail = newNode
	}
}

func (ll *LinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{
		value: value,
	}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}
}

func (ll *LinkedList[T]) Delete(value T) {
	if ll.head == nil {
		return
	}
	if ll.head.value == value {
		ll.head = ll.head.next
		if ll.head == nil {
			ll.tail = nil
		}
		return
	}

	for current := ll.head; current.next != nil; current = current.next {
		if current.next.value == value {
			current.next = current.next.next
			if current.next == nil {
				ll.tail = current
			}
			return
		}
	}
}

func (ll *LinkedList[T]) Contains(value T) bool {
	for current := ll.head; current != nil; current = current.next {
		if current.value == value {
			return true
		}
	}

	return false
}

func (ll *LinkedList[T]) Display() {
	for current := ll.head; current != nil; current = current.next {
		fmt.Printf("%v -> ", current.value)
	}
	fmt.Println("nil")
}
