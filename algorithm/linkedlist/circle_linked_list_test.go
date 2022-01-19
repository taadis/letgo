package linkedlist

import "testing"

func TestCircleLinkedList_InsertAfter(t *testing.T) {
	list := NewCircleLinkedList()
	list.Append(1)
	list.InsertAfter(0, 3)
	list.InsertAfter(1, 5)
	list.InsertAfter(2, 7)
	list.Shows()
	t.Logf("length:%d", list.Length())
}

func TestCircleLinkedList_InsertBefore(t *testing.T) {
	list := NewCircleLinkedList()
	list.Append(1)
	list.InsertBefore(0, 3)
	list.InsertBefore(0, 5)
	list.InsertBefore(0, 7)
	list.Shows()
}

func TestCircleLinkedList_Remove(t *testing.T) {
	list := NewCircleLinkedList()
	list.Append(1)
	list.Append(3)
	list.Append(5)
	list.Shows()

	list.Remove(2)
	list.Remove(1)
	list.Remove(0)
	list.Shows()
}

func TestCircleLinkedList_IsEmpty(t *testing.T) {
	list := NewCircleLinkedList()
	list.Shows()
	t.Logf("IsEmpty:%v", list.IsEmpty())
}

func TestCircleLinkedList_Append(t *testing.T) {
	list := NewCircleLinkedList()
	t.Logf("IsEmpty:%v", list.IsEmpty())

	list.Append(1)
	list.Append(3)
	list.Append(5)
	t.Logf("IsEmpty:%v", list.IsEmpty())
	t.Logf("Length:%d", list.Length())
	list.Shows()
}
