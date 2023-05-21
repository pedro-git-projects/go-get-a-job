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
func ProcessElements(n *html.Node, selector string, callback NodeProcessor) {
	processNode := *new(NodeProcessor)
	processNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == selector {
			callback(n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			processNode(c)
		}
	}
	processNode(n)
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
