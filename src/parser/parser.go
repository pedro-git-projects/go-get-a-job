package parser

import (
	"strings"

	"golang.org/x/net/html"
)

// NodeProcessor declares the callback signature for processing nodes
type NodeProcessor func(*html.Node)

// ProcessElements traverses the HTML node tree starting from the given root node (n),
// filters elements based on the provided selector string, and invokes the callback
// function (callback) for each matching element. It recursively explores the node tree,
// processing elements that match the selector and their descendants.
//
// The callback function (callback) takes an *html.Node as a parameter and is responsible
// for processing the matching elements according to specific requirements.
//
// The traversal is performed using a stack-based iteration to avoid potential stack overflows
// when processing large HTML trees. The function creates a stack and pushes the root node onto it.
// It then repeatedly pops a node from the stack, checks if it matches the selector,
// and invokes the callback function if it does. The function also pushes all the node's children
// onto the stack for further traversal.
//
// Note that the callback function should handle the processing of individual elements according
// to specific requirements, while this function focuses on traversing and filtering the HTML node tree.
func ProcessElements(n *html.Node, selector string, callback NodeProcessor) {
	stack := &nodeStack{}
	stack.push(n)

	for {
		node := stack.pop()
		if node == nil {
			break
		}

		if node.Type == html.ElementNode && node.Data == selector {
			callback(node)
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			stack.push(c)
		}
	}
}

// GetText extracts the text content from the given HTML node (n) and returns it.
// If the node is a text node, the data within it is returned. If the node is not
// a text node, the function recursively explores the node's descendants, concatenates
// their text contents, and returns the result.
//
// The returned text is trimmed of leading and trailing whitespace.
func GetText(n *html.Node) (text string) {
	if n.Type == html.TextNode {
		text = n.Data
	} else {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			text += GetText(c)
		}
	}
	return strings.TrimSpace(text)
}
