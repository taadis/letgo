package linkedlist

import (
	"fmt"
	"sync"
)

type DoubleLinkedLister interface {
	RemoveAt(index int)
	Append(data interface{})
	IsEmpty() bool
	Length() int
	Get(index int) *DoubleLinkedNode
	Shows()
}

type DoubleLinkedNode struct {
	Prior *DoubleLinkedNode // 前驱节点指针域
	Data  interface{}       // 数据域
	Next  *DoubleLinkedNode // 后继节点指针域
}

type DoubleLinkedList struct {
	head   *DoubleLinkedNode
	length int
	mutex  *sync.RWMutex
}

func NewDoubleLinkedList() DoubleLinkedLister {
	l := new(DoubleLinkedList)
	l.head = new(DoubleLinkedNode)
	l.length = 0
	l.mutex = new(sync.RWMutex)
	return l
}

func (l *DoubleLinkedList) RemoveAt(index int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if index < 0 || index > l.Length()-1 {
		return
	}

	p := l.head
	for j := 0; j <= index; j++ {
		p = p.Next
	}
	p.Prior.Next = p.Next
	if p.Next != nil {
		p.Next.Prior = p.Prior
	}
	l.length--
}

func (l *DoubleLinkedList) Append(data interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// new node
	node := new(DoubleLinkedNode)
	node.Data = data

	p := l.head
	for p.Next != nil {
		p = p.Next
	}

	node.Prior = p
	p.Next = node
	l.length++
}

func (l *DoubleLinkedList) Shows() {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	if l.IsEmpty() {
		return
	}

	p := l.head.Next
	for p != nil {
		fmt.Printf("\tDoubleLinkedNode data:%v\n", p.Data)
		p = p.Next
	}
}

func (l *DoubleLinkedList) IsEmpty() bool {
	if l.head.Next == nil {
		return true
	}
	return false
}

func (l *DoubleLinkedList) Length() int {
	return l.length
}

func (l *DoubleLinkedList) Get(index int) *DoubleLinkedNode {
	if index < 0 || index > l.Length()-1 {
		return nil
	}

	p := l.head
	for j := 0; j <= index; j++ {
		p = p.Next
	}
	return p
}
