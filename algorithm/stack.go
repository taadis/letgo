package algorithm

import "fmt"

type Node struct {
	Value int
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

type Stack struct {
	nodes []*Node
	count int
}

// NewStack returns a new stack
func NewStack() *Stack {
	return &Stack{}
}

// Push adds a node to the top of the stack
func (s *Stack) Push(node *Node) {
	s.nodes = append(s.nodes, node)
	s.count++
}

// Pop removes and returns a node from the top of the stack
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}
