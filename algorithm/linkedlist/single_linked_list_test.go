package linkedlist

import (
	"testing"
)

func TestSingleLinkedList_Remove(t *testing.T) {
	list := NewSingeLinkedList()
	list.Append(1)
	list.Append(3)
	list.Append(5)
	list.Shows()
	list.RemoveAt(1)
	list.Shows()
}

// 测试单向链表
func TestSingleLinkedList(t *testing.T) {
	// test1
	list := NewSingeLinkedList()
	list.Shows()

	// test2
	list.Append("3")
	list.Append("5")
	list.Append("7")
	list.Shows()
	t.Logf("test2 length:%d", list.Length())

	// test3
	list.Prepend("1")
	list.Shows()
	t.Logf("test3 length:%d", list.Length())

	// test4
	list.Insert(3, "0")
	list.Shows()
}


