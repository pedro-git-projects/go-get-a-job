package parser

import "golang.org/x/net/html"

// nodeStack is used to provide stack-based iterations instead of recursion,
// particularly in ProcessElements, in order to avoid potential stack overflows
// when processing large HTML trees.
type nodeStack struct {
	nodes []*html.Node
}

// push adds a new HTML node to the top of the stack.
func (s *nodeStack) push(n *html.Node) {
	s.nodes = append(s.nodes, n)
}

// pop removes and returns the HTML node at the top of the stack.
// If the stack is empty, it returns nil.
func (s *nodeStack) pop() *html.Node {
	if len(s.nodes) == 0 {
		return nil
	}
	node := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return node
}
