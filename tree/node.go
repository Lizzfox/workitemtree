package tree

import (
	"fmt"
)

type Node struct {
	ID       int
	URL      string
	Title    string
	Children []*Node
}

func (node *Node) show(prefix string) {
	if prefix == "" {
		fmt.Printf("|-%v %v\n", node.ID, node.Title)
	} else {
		fmt.Printf("%v |-%v %v\n", prefix, node.ID, node.Title)
	}

	for _, n := range node.Children {
		n.show(prefix+"  ")
	}
}
