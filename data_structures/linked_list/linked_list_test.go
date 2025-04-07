package linked_list

import (
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	ll := New[int]()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	if !ll.Contains(1) || !ll.Contains(2) || !ll.Contains(3) {
		t.Errorf("Append failed: expected elements to be present")
	}

	if ll.head.value != 1 || ll.tail.value != 3 {
		t.Errorf("Head/Tail not correct after Append")
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	ll := New[int]()
	ll.Prepend(1)
	ll.Prepend(2)
	ll.Prepend(3)

	if !ll.Contains(1) || !ll.Contains(2) || !ll.Contains(3) {
		t.Errorf("Prepend failed: expected elements to be present")
	}

	if ll.head.value != 3 || ll.tail.value != 1 {
		t.Errorf("Head/Tail not correct after Prepend")
	}
}

func TestLinkedList_Delete(t *testing.T) {
	ll := New[int]()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)

	ll.Delete(1) // delete head
	if ll.head.value != 2 {
		t.Errorf("Delete failed: head not updated after deleting 1")
	}

	ll.Delete(4) // delete tail
	if ll.tail.value != 3 {
		t.Errorf("Delete failed: tail not updated after deleting 4")
	}

	ll.Delete(3) // delete middle
	if ll.Contains(3) {
		t.Errorf("Delete failed: value 3 still present")
	}

	ll.Delete(999) // non-existent value
	// should not crash or modify the list
}

func TestLinkedList_Contains(t *testing.T) {
	ll := New[string]()
	ll.Append("a")
	ll.Append("b")

	if !ll.Contains("a") || !ll.Contains("b") {
		t.Errorf("Contains failed: expected values not found")
	}

	if ll.Contains("c") {
		t.Errorf("Contains failed: unexpected value found")
	}
}

func TestLinkedList_Empty(t *testing.T) {
	ll := New[int]()

	if ll.Contains(42) {
		t.Errorf("Contains should return false on empty list")
	}

	ll.Delete(42) // should not panic
}
