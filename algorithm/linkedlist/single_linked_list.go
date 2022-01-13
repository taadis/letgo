package linkedlist

import (
	"errors"
	"fmt"
	"sync"
)

// 单向链表的操作方法集
type SingleLinkedLister interface {
	Length() int                        // 获取长度
	Prepend(data interface{})           // 前置追加
	Append(data interface{})            // 后置追加
	Insert(index int, data interface{}) // 在指定index位置插入节点
	RemoveAt(index int) error           // 移除指定index位置的节点
	Shows()                             // 显示所有
}

// 节点数据
type DataType interface{}

// 单向链表的节点
type SingleLinkedNode struct {
	Data DataType          // 数据域
	Next *SingleLinkedNode // 指针域,指向下一个节点的地址
}

// 单向链表
type SingleLinkedList struct {
	mutex  *sync.RWMutex
	Head   *SingleLinkedNode
	length int
}

// 单向链表的初始化,带独立头节点,以便后续操作
func NewSingeLinkedList() SingleLinkedLister {
	l := new(SingleLinkedList)
	l.mutex = new(sync.RWMutex)
	l.Head = new(SingleLinkedNode)
	l.length = 0
	return l
}

func (l *SingleLinkedList) RemoveAt(index int) error {
	if index < 0 || index > l.length-1 {
		return errors.New("invalid index:%d when remove")
	}

	if l.IsEmpty() {
		return nil
	}

	pre := l.Head
	p := l.Head
	for j := 0; j <= index; j++ {
		pre = p
		p = p.Next
	}

	pre.Next = p.Next
	return nil
}

func (l *SingleLinkedList) Insert(index int, data interface{}) {
	if index < 0 {
		l.Prepend(data)
		return
	}

	if index > l.length-1 {
		l.Append(data)
	}

	// 定位需要插入的节点
	p := l.Head
	for i := 0; i < index; i++ {
		p = p.Next
	}

	// 创建新的节点
	node := new(SingleLinkedNode)
	node.Data = data
	node.Next = p.Next
	p.Next = node
	l.length++
}

func (l *SingleLinkedList) Length() int {
	return l.length
}

func (l *SingleLinkedList) Prepend(data interface{}) {
	// 创建新的节点
	node := new(SingleLinkedNode)
	node.Data = data
	node.Next = l.Head.Next
	l.Head.Next = node
	l.length++
}

// 尾部追加
func (l *SingleLinkedList) Append(data interface{}) {
	// 新创建的节点
	node := new(SingleLinkedNode)
	node.Data = data

	// 如果链表不为空
	p := l.Head
	for p.Next != nil {
		p = p.Next
	}

	// 此时p为尾节点next,将其指向新创建的节点
	p.Next = node
	l.length++
}

// 遍历所有节点
func (l *SingleLinkedList) Shows() {
	if l.IsEmpty() {
		return
	}

	p := l.Head
	for p.Next != nil {
		p = p.Next
		fmt.Printf("\t单向链表数据:%v\n", p.Data)
	}
}

// 判断单向链表是否为空
func (l *SingleLinkedList) IsEmpty() bool {
	// 判断单向链表是否为空,只需判断头节点是否为空即可
	if l.Head.Next == nil {
		return true
	}
	return false
}
