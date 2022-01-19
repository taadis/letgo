package linkedlist

import "testing"

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
