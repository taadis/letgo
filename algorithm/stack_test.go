package algorithm

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack()
	node := stack.Pop()
	t.Log(node)
	stack.Push(&Node{Value: 1})
	stack.Push(&Node{Value: 3})
	stack.Push(&Node{Value: 5})
	t.Log(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
}
