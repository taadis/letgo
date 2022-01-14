package linkedlist

import "testing"

func TestDoubleLinkedList_Append(t *testing.T) {
	list := NewDoubleLinkedList()
	list.Append(1)
	list.Append(3)
	list.Append(5)
	list.Shows()
	t.Logf("len:%d", list.Length())

	list.RemoveAt(2)
	list.RemoveAt(1)
	list.RemoveAt(0)
	list.Shows()
	t.Logf("len:%d", list.Length())

	list.Append(2)
	list.Append(4)
	list.Append(6)
	list.Shows()
	t.Logf("len:%d", list.Length())
}

func TestDoubleLinkedList_Shows(t *testing.T) {
	list := NewDoubleLinkedList()
	list.Shows()
}
