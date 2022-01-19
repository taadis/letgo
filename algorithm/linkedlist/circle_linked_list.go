package linkedlist

import (
	"fmt"
	"sync"
)

type CircleLinkedLister interface {
	IsEmpty() bool
	Length() int
	Append(data interface{})                  // 尾部追加节点
	InsertBefore(index int, data interface{}) // 在指定索引位置之前插入
	InsertAfter(index int, data interface{})  // 在指定索引位置之后插入
	Remove(index int)                         // 移除指定索引位置的节点
	Shows()                                   // 遍历显示
}

type CircleLinkedNode struct {
	//Prior *CircleLinkedNode // 指向前驱节点的指针域
	Data interface{}       // 数据域
	Next *CircleLinkedNode // 指向后继节点的指针域
}

type CircleLinkedList struct {
	head   *CircleLinkedNode // 头指针
	length int               // 链表长度
	mutex  *sync.RWMutex     // 读写锁
}

func NewCircleLinkedList() CircleLinkedLister {
	l := new(CircleLinkedList)
	head := &CircleLinkedNode{}
	head.Next = head
	l.head = head
	l.length = 0
	l.mutex = new(sync.RWMutex)
	return l
}

func (l *CircleLinkedList) IsEmpty() bool {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	if l.head.Next == l.head {
		return true
	}
	return false
}

func (l *CircleLinkedList) Length() int {
	return l.length
}

func (l *CircleLinkedList) Append(data interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// 新创建的节点
	node := new(CircleLinkedNode)
	node.Data = data   // 新节点数据域赋值
	node.Next = l.head // 追加的节点,next默认指向头节点

	// 找到尾节点,即指针域next为head的节点
	p := l.head
	for p.Next != l.head {
		p = p.Next
	}
	p.Next = node
	l.length++
}

func (l *CircleLinkedList) InsertBefore(index int, data interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if index < 0 || index > l.Length()-1 {
		return
	}

	// new node
	node := &CircleLinkedNode{
		Data: data,
		Next: nil,
	}

	p := l.head // 指定索引位置的节点
	q := l.head // 前一个节点
	for j := 0; j <= index; j++ {
		q = p
		p = p.Next
	}

	node.Next = p
	q.Next = node
	l.length++
}

func (l *CircleLinkedList) InsertAfter(index int, data interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if index < 0 || index > l.Length()-1 {
		return
	}

	// find index node
	p := l.head
	for j := 0; j <= index; j++ {
		p = p.Next
	}

	// new node
	node := &CircleLinkedNode{
		Data: data,
		Next: p.Next,
	}
	p.Next = node
	l.length++
}

func (l *CircleLinkedList) Remove(index int) {
	if index < 0 || index > l.Length()-1 {
		return
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	// 找到指定节点
	p := l.head
	q := l.head
	for j := 0; j <= index; j++ {
		q = p
		p = p.Next
	}
	q.Next = p.Next
	l.length--
}

func (l *CircleLinkedList) Shows() {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	if l.IsEmpty() {
		fmt.Printf("is empty when shows\n")
	}

	p := l.head
	for p.Next != l.head {
		p = p.Next
		fmt.Printf("node.data:%v\n", p.Data)
	}
}
