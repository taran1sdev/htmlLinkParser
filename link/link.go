package link

import (
	"io"
	"golang.org/x/net/html"
)

type Anchor struct{
	Href string
	Text string
}

// anchorNodes recursively searches through the html tree checking if ElementNodes are Anchor tags

func anchorNodes(n *html.Node) (nodes []*html.Node) {
	// Check if node is an <a> tag and if it is return it
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	
	// dfs search through html tree appending each node we find
	// to our return variable
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, anchorNodes(c)...)
	} 

	return
}

func ParseAnchors(r io.Reader) (anchors []Anchor, err error) {
	// Get the first node in the html tree
	docNode, err := html.Parse(r)
	if err != nil {
		return nil, err	
	}
	
	// recursively search the tree for <a> nodes and return them to variable nodes
	nodes := anchorNodes(docNode)

	// print out the <a> nodes
	for _, node := range nodes {
		// make sure the node has a href and if so create an anchor variable with the
		// href and text of the first child then append it to the return value
		if node.Attr[0].Key == "href" {
			a := Anchor{node.Attr[0].Val, node.FirstChild.Data}
			anchors = append(anchors, []Anchor{a}...)
		}
	}
	return 
}
