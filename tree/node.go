package tree

import (
	"fmt"
)

type Node struct {
	ID int
	URL string
	Children []*Node
}


func showNode(node *Node, prefix string){
	if prefix == "" {
		fmt.Printf("")
	}
}