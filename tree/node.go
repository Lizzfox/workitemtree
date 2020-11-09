package tree

import (
	"fmt"
)

// Node struct represents a single unit of work item tree 
type Node struct {
	ID       int
	URL      string
	Title    string
	Children []*Node
}

func (node *Node) show(prefix string) {
	resStr :=  fmt.Sprintf("%v %v\n", node.ID, node.Title)

	if prefix != "" {
		resStr = prefix + resStr
	}

	fmt.Print(resStr)

	for _, n := range node.Children {
		n.show(prefix + "  ")
	}
}
